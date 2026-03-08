pipeline {
    agent any

    environment {
        APP_NAME = "chatbot-pendidikan-tinggi"
        DOCKER_IMAGE = "peenesss/chatbot-pendidikan-tinggi"
        DOCKER_TAG = "latest"
        APP_PORT = "9091"

        GO_BIN = "/usr/local/go/bin/go"
        PATH = "/usr/local/go/bin:${env.PATH}"

        // DockerHub credential
        DOCKER_CREDS = "dockerhub-creds"

        // Database
        DB_HOST = "103.149.177.39"
        DB_USER = "johannessipayung"
        DB_PASSWORD = "password123"
        DB_NAME = "auth_golang_clean"
        DB_PORT = "5432"
        DB_SSLMODE = "disable"

        // VPS
        VPS_USER = "root"
        VPS_HOST = "103.149.177.39"
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
                sh '''
                docker build -t ${DOCKER_IMAGE}:${DOCKER_TAG} -f docker/Dockerfile .
                '''
            }
        }

        stage('Login Docker Hub') {
            steps {
                withCredentials([usernamePassword(
                    credentialsId: "${DOCKER_CREDS}",
                    usernameVariable: 'DOCKER_USER',
                    passwordVariable: 'DOCKER_PASS'
                )]) {

                    sh '''
                    echo $DOCKER_PASS | docker login -u $DOCKER_USER --password-stdin
                    '''
                }
            }
        }

        stage('Push Image to Docker Hub') {
            steps {
                sh '''
                docker push ${DOCKER_IMAGE}:${DOCKER_TAG}
                '''
            }
        }

        stage('Deploy to VPS') {
            steps {
                sshagent(['vps-ssh']) {

                    sh """
                    ssh -o StrictHostKeyChecking=no ${VPS_USER}@${VPS_HOST} '

                        echo "Pull latest image from Docker Hub..."
                        docker pull ${DOCKER_IMAGE}:${DOCKER_TAG}

                        echo "Stop old container..."
                        docker stop ${APP_NAME} || true
                        docker rm ${APP_NAME} || true

                        echo "Run new container..."
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
                        ${DOCKER_IMAGE}:${DOCKER_TAG}

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