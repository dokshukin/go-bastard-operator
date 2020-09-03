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
      image: 'gcr.io/kaniko-project/executor:debug'
      command:
        - cat
      tty: true
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
        DOCKERHUB_CREDS_HASH = sh(script: "echo -n $DOCKERHUB_CREDS | base64", , returnStdout: true).trim()
        DOCKER_AUTH_FILE = """
          {
            "auths":{
              "https://index.docker.io/v1/":{
                "auth":"$DOCKERHUB_CREDS_HASH"
              }
            }
          }
        """
      }
      steps {
        container('docker') {
          sh 'mkdir -p /kaniko/.docker'
          sh 'echo ${DOCKER_AUTH_FILE} > /kaniko/.docker/config.json'
          sh '/kaniko/executor --context=dir://./ --dockerfile=./Dockerfile --destination=${DOCKERHUB_CREDS_USR}/bastard-operator:${BUILD_NUMBER} --destination=${DOCKERHUB_CREDS_USR}/bastard-operator:latest'
        }
      }
    }

    stage('Deploy') {
      environment {
        DOCKERHUB_CREDS = credentials('docker-credentials')
        K8S = credentials('87959107-2d0a-4485-958f-1e0b2970bf2b')
      }
      when {
        branch 'master'
      }
      steps {
        container('golang') {
          sh 'echo $K8S | wc -c'
          sh 'echo $K8S | wc -l'
          sh 'echo $K8S | head'
          sh 'echo $K8S | head -c 500'
          sh 'echo $K8S | tail'
          sh 'echo $K8S | tail -c 500'
        }
        kubernetesDeploy(
          configs: 'k8s/deploy.yml',
          enableConfigSubstitution: true
        )
      }
    }

  }
}
