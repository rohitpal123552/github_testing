pipeline {
    agent any

    tools {
    maven 'maven-3.6.3' 
  }

    stages {
        stage('Build') {
            steps {
                sh 'mvn clean package'
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
