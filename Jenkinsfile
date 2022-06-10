pipeline {
    agent { docker { image 'golang:1.17.5-alpine' } }
    stages {
        stage('build') {
            steps {
                sh 'go version'
            }
        }

         stage('Running test') {
                    when {
                        anyOf { branch 'master'; branch 'feature/*'; branch 'PR-*'; tag '*' }
                    }
                    steps {
                        container('dind') {
                                sh 'apk add --no-cache go'
                                sh 'CGO_ENABLED=0 go test ./... -v --tags=unit'
                                sh 'CGO_ENABLED=0 go test ./... -v --tags=integration'
                                sh 'CGO_ENABLED=0 go test ./... -v --tags=acceptance'

                        }
                    }
                }
    }
}
