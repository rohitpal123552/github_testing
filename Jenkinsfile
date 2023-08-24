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
                println(GOPATH)
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
                sh 'go test ./...'
            }
        }
        stage('Run') {
            steps {
                echo 'Start the Go application'
                sh './jenkins_go_project'
            }
        }
        stage('Create Tarball') {
            steps {
                sh 'tar -czvf jenkins_go_project.tar.gz jenkins_go_project'
                script {
                    // Define the filename
                    def filename = 'jenkins_go_project.tar.gz'
                    
                    // Check if the tarball file exists
                    def tarballExists = fileExists(filename)
                    if (tarballExists) {
                        def currentDir = pwd()
                        def fullPath = "${currentDir}/${filename}"
                        echo "Tarball created at: ${fullPath}"
                    } else {
                        error 'Tarball creation failed'
                    }
                }
            }
        }
        // stage('Create Zip') {
        //     steps {
        //         echo 'Create a zip file containing the built binary'
        //         sh 'zip jenkins_go_project.zip jenkins_go_project'
        //     }
        // }
        
    }
}
