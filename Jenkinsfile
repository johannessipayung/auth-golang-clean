pipeline {
    agent any

    environment {
        APP_NAME = "chatbot-pendidikan-tinggi"
        DOCKER_IMAGE = "chatbot-pendidikan-tinggi:latest"
        PORT = "9090"
        HOST_PORT = "9090"      // Port yang diakses dari luar (macOS)
        CONTAINER_PORT = "8080" // Port yang ditulis di main.go
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
                go test ./tests -v -coverprofile=coverage.out
                '''
            }
        }
        stage('Show Coverage') {
            steps {
                sh '''
                if [ -f coverage.out ]; then
                    go tool cover -func=coverage.out
                fi
                '''
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
                sh 'docker build -t chatbot-pendidikan-tinggi:latest -f docker/Dockerfile .'
            }
        }

        stage('Run Container Test') {
            steps {
                sh '''
                docker rm -f chatbot-test || true
                
                # Map HOST_PORT (9090) ke CONTAINER_PORT (8080)
                docker run -d -p $HOST_PORT:$CONTAINER_PORT --name chatbot-test $DOCKER_IMAGE
                
                sleep 10
                
                # Test ke port 9090
                curl -f http://localhost:$HOST_PORT || (docker logs chatbot-test && exit 1)
                
                docker rm -f chatbot-test
                '''
            }
        }

        stage('Deploy Container') {
            steps {
                sh '''
                docker stop $APP_NAME || true
                docker rm $APP_NAME || true

                docker run -d \
                -p $HOST_PORT:$CONTAINER_PORT \
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