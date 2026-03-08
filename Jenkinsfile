pipeline {
    agent any

    environment {
        APP_NAME = "chatbot-pendidikan-tinggi"
        DOCKER_IMAGE = "chatbot-pendidikan-tinggi:latest"
        APP_PORT = "9090"
        GO_BIN = "/usr/local/go/bin/go"
        PATH = "/usr/local/go/bin:${env.PATH}"

        // Database environment (VPS)
        DB_HOST = "103.149.177.39"           // IP VPS
        DB_USER = "johannessipayung"
        DB_PASSWORD = "password123"
        DB_NAME = "auth_golang_clean"
        DB_PORT = "5432"
        DB_SSLMODE = "disable"
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
                ${GO_BIN} version
                ${GO_BIN} mod tidy
                ${GO_BIN} mod download
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

        stage('Quality Checks') {
            parallel {
                stage('Lint') {
                    steps {
                        sh './bin/golangci-lint run ./...'
                    }
                }

                stage('Vet') {
                    steps {
                        sh '${GO_BIN} vet ./...'
                    }
                }

                stage('Unit Tests') {
                    steps {
                        sh '${GO_BIN} test -v -coverprofile=coverage.out ./tests'
                    }
                }
            }
        }

        stage('Build Application') {
            steps {
                sh '${GO_BIN} build -o app ./cmd'
            }
        }

        stage('Build Docker Image') {
            steps {
                sh 'docker build -t ${DOCKER_IMAGE} -f docker/Dockerfile .'
            }
        }

        stage('Deploy to VPS') {
            steps {
                sshagent(['vps-ssh']) {   // Credential SSH di Jenkins
                    sh """
                    ssh -o StrictHostKeyChecking=no root@${DB_HOST} '
                        echo "Stopping old container if exists..."
                        docker stop ${APP_NAME} || true
                        docker rm ${APP_NAME} || true

                        echo "Running new container..."
                        docker run -d \
                            --name ${APP_NAME} \
                            -p ${APP_PORT}:8080 \
                            --restart unless-stopped \
                            -e DB_HOST=${DB_HOST} \
                            -e DB_USER=${DB_USER} \
                            -e DB_PASSWORD=${DB_PASSWORD} \
                            -e DB_NAME=${DB_NAME} \
                            -e DB_PORT=${DB_PORT} \
                            -e DB_SSLMODE=${DB_SSLMODE} \
                            ${DOCKER_IMAGE}
                    '
                    """
                }
            }
        }

        stage('Archive Coverage') {
            steps {
                archiveArtifacts artifacts: 'coverage.out', fingerprint: true
            }
        }

        stage('Debug') {
            steps {
                sh 'docker --version'
                sh 'which go'
                sh 'docker ps -a'
            }
        }
    }

    post {
        always {
            echo 'Pipeline finished'
        }
        success {
            echo 'CI/CD Pipeline SUCCESS'
        }
        failure {
            echo 'CI/CD Pipeline FAILED'
        }
    }
}