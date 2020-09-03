pipeline {
  agent {
    kubernetes {
      yaml """
apiVersion: v1
kind: Pod
metadata:
  labels:
    label: runner
spec:
  containers:

    - name: golang
      image: 'mjslabs/jenkins-jnlp-golang:latest'
      command:
        - cat
      tty: true

    - name: docker
      image: docker:18.05-dind
      securityContext:
        privileged: true
      volumeMounts:
        - name: dind-storage
          mountPath: /var/lib/docker

    - name: kubectl
      image: 'yonadev/jnlp-slave-k8s-helm:latest'
      command:
        - cat
      tty: true

    - name: envsubst
      image: 'cirocosta/alpine-envsubst:latest'
      command:
        - cat
      tty: true

  volumes:
    - name: dind-storage
      emptyDir: {}
"""
    }
  }


  stages {

    stage('Tests') {
      steps {
        container('golang') {
          sh 'go version'
          sh 'go test'
        }
      }
    }

    stage('Build binary') {
      environment {
        GOOS = 'linux'
        GOARCH = 'amd64'
        TIMESTAMP = sh(script: '/bin/date -u +%Y-%m-%dT%T%z', , returnStdout: true).trim()
      }
      steps {
        container('golang') {
          sh "go build --ldflags='-X \"main.Version=${BUILD_NUMBER}\" -X \"main.Timestamp=${TIMESTAMP}\"' -o bastard_operator ."
        }
      }
    }

    stage('Build and publish docker image') {
      when {
        branch 'master'
      }
      environment {
        DOCKERHUB_CREDS = credentials('docker-credentials')
      }
      steps {
        container('docker') {
          sh 'docker login -u ${DOCKERHUB_CREDS_USR} -p ${DOCKERHUB_CREDS_PSW}'
          sh 'docker build -t ${DOCKERHUB_CREDS_USR}/bastard-operator:${BUILD_NUMBER} -t ${DOCKERHUB_CREDS_USR}/bastard-operator:latest .'
          sh 'docker push ${DOCKERHUB_CREDS_USR}/bastard-operator:${BUILD_NUMBER}'
          sh 'docker push ${DOCKERHUB_CREDS_USR}/bastard-operator:latest'
        }
      }
    }

    stage('Deploy') {
      environment {
        DOCKERHUB_CREDS = credentials('docker-credentials')
        KUBE = credentials('k8s-service-account')
      }
      when {
        branch 'master'
      }
      steps {
        container('envsubst') {
          sh 'envsubst < k8s/deploy.yml > k8s/deploy_prepared.yml'
          sh 'cat k8s/deploy_prepared.yml'
        }
        container('kubectl') {
          sh 'mkdir ~/.kube'
          sh 'echo "$KUBE" | base64 -d > ~/.kube/config'
          sh 'kubectl -n jenkins -f k8s/deploy_prepared.yml apply'
        }
      }
    }

  }
}
