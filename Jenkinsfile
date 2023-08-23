pipeline {
    agent any
    tools {
        go 'GO-1.20'
    }
    environment {
        GO114MODULE = 'on'
        CGO_ENABLED = 0 
        GOPATH = "${JENKINS_HOME}/jobs/${JOB_NAME}/builds/${BUILD_ID}"
    }
    stages {
        stage('Checkout') {
            steps {
                echo "Checkout your Go project from the version control system (e.g., Git)"
                checkout scm
            }
        }
        stage('Pre Test') {
            steps {
                echo 'Installing dependencies'
                sh 'go version'
            }
        }
        stage('Build') {
            steps {
                echo 'Build the Go project'
                sh 'go build'
            }
        }
        stage('Test') {
            steps {
                echo 'Run unit tests'
                sh 'go test ./mymath'
            }
        }
        stage('Run') {
            steps {
                echo 'Start the Go application'
                sh './jenkins_go_project'
            }
        }
        
    }
}
