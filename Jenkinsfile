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
        stage('Pre Test') {
            steps {
                echo 'Installing dependencies'
                echo "${GOPATH}"
                sh 'go version'
                sh 'go get -u golang.org/x/lint/golint'
            }
        }
        
        stage('Build') {
            steps {
                echo 'Compiling and building'
                sh 'go build'
                // sh './go-project-pipeline'
            }
        }

        // stage('Test') {
        //     steps {
        //         withEnv(["PATH+GO=${GOPATH}/bin"]){
        //             echo 'Running vetting'
        //             sh 'go vet .'
        //             echo 'Running linting'
        //             sh 'golint .'
        //             echo 'Running test'
        //             sh 'cd test && go test -v'
        //         }
        //     }
        // }
        
    }
}
