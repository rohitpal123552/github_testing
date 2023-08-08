pipeline {
    agent any

    stages {
        stage('Build') {
            steps {
                echo 'Building the application...'
                script {
                    def test = 2+2 > 3 ? "cool" : "not cool"
                    echo test
                }
            }
        }
        stage('Test') {
            steps {
                echo 'Testing..'
            }
        }
        stage('Deploy') {
            steps {
                echo 'Deploying....'
            }
        }
    }
}
