pipeline {
    agent {
        docker { image 'golang:1.18-stretch' }
    }
    environment {
        GO114MODULE = 'on'
        CGO_ENABLED = 0 
        GOPATH = "/go"
        HOME = "/home/perolo/Jenkins/workspace/${JOB_NAME}"
    }
    stages {        
        stage('Pre Test') {
            steps {
                echo 'Installing dependencies'
                sh 'env'
                sh 'pwd'             
                sh 'go version'
                sh 'go install honnef.co/go/tools/cmd/staticcheck@latest'
                sh 'go install github.com/jstemmer/go-junit-report@latest'
                sh 'go install github.com/axw/gocov/gocov@latest'
                sh 'go install github.com/AlekSi/gocov-xml@latest'
                sh 'go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest'
            }
        }
        
        stage('Build') {
            steps {
                echo 'Compiling and building'
                sh 'go build'
            }
        }

        stage('Test') {
            steps {
                withEnv(["PATH+GO=${GOPATH}/bin"]){
                    catchError(buildResult: 'SUCCESS', stageResult: 'UNSTABLE', message: 'Static codecheck errors!') {
                        echo 'Running vetting'
                        sh 'go vet .'
                        echo 'Running staticcheck'
                        sh 'staticcheck ./...'
                        echo 'Running golangci-lint'
                        sh 'golangci-lint run --out-format junit-xml --config .golangci.yml > golangci-lint.xml'
                    }
                    echo 'Running test'
                    sh 'go test -v 2>&1 | go-junit-report > report.xml'
                    echo 'Running coverage'
                    sh 'gocov test ./... | gocov-xml > coverage.xml'
                }
            }
        }
        stage('Artifacts') {
            steps { 
                script {
                    if (fileExists('report.xml')) {
                        archiveArtifacts artifacts: 'report.xml', fingerprint: true
                        junit 'report.xml'
                    }
                    if (fileExists('coverage.xml')) {
                        archiveArtifacts artifacts: 'coverage.xml', fingerprint: true
                        cobertura coberturaReportFile: 'coverage.xml'
                    }
                    if (fileExists('golangci-lint.xml')) {
                        archiveArtifacts artifacts: 'golangci-lint.xml'            
                        try {
                            junit 'golangci-lint.xml'
                        } catch (err) {
                            echo err.getMessage()
                            echo "Error detected, but we will continue."
                            echo "No lint errors found is not an error."
                        }
                    }
                }
            }   
        }
    }
}
