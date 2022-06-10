pipeline {
    agent { docker { image '1.18.3-bullseye' } }
    stages {
        stage('build') {
            steps {
                sh 'go version'
            }
        }
        stage('SonarQube Analysis') {
            def scannerHome = tool 'sonarqube-local';
            withSonarQubeEnv() {
              sh "${scannerHome}/bin/sonar-scanner"
            }
          }
    }
}
