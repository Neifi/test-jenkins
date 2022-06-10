pipeline {
    agent { docker { image '1.18.3-bullseye' } }
    stages {
        stage('build') {
            steps {
                sh 'go version'
            }
        }
    }
}
