pipeline {
    agent any

    environment {
        APP_NAME = "chatbot-pendidikan-tinggi"
        DOCKER_IMAGE = "chatbot-pendidikan-tinggi:latest"
        PORT = "9090"
    }

    stages {

        stage('Checkout Code') {
            steps {
                git branch: 'main', url: 'https://github.com/johannessipayung/auth-golang-clean.git'
            }
        }

        stage('Setup Go Modules') {
            steps {
                sh '''
                go version
                go mod tidy
                go mod download
                '''
            }
        }

        stage('Install golangci-lint') {
            steps {
                sh '''
                curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh \
                | sh -s latest
                '''
            }
        }

        stage('Run Linter') {
            steps {
                sh './bin/golangci-lint run ./...'
            }
        }

        stage('Go Vet') {
            steps {
                sh 'go vet ./...'
            }
        }

        stage('Run Unit Tests') {
            steps {
                sh '''
                go test ./... -v -coverprofile=coverage.out || exit 1
                '''
            }
        }
        stage('Show Coverage') {
            steps {
                sh 'go tool cover -func=coverage.out'
            }
        }

        stage('Build Application') {
            steps {
                sh '''
                go build -o app ./cmd
                '''
            }
        }

        stage('Build Docker Image') {
            steps {
                sh '''
                docker build -t $DOCKER_IMAGE .
                '''
            }
        }

        stage('Run Container Test') {
            steps {
                sh '''
                docker run -d -p $PORT:$PORT --name chatbot-test $DOCKER_IMAGE

                sleep 10

                curl -f http://localhost:$PORT || exit 1

                docker stop chatbot-test
                docker rm chatbot-test
                '''
            }
        }

        stage('Deploy Container') {
            steps {
                sh '''
                docker stop $APP_NAME || true
                docker rm $APP_NAME || true

                docker run -d \
                -p $PORT:$PORT \
                --name $APP_NAME \
                $DOCKER_IMAGE
                '''
            }
        }
    }

    post {

        always {
            archiveArtifacts artifacts: 'coverage.out', fingerprint: true
        }

        success {
            echo 'CI/CD Pipeline SUCCESS'
        }

        failure {
            echo 'CI/CD Pipeline FAILED'
        }
    }
}