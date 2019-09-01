module.exports = {
    apps : [{
      name: "lingotalk-exam-back",
      script: "./back",
      env: {
        GO_ENV: "staging",
      },
      env_development: {
        GO_ENV: "development",
      },
      env_development_am: {
        GO_ENV: "development",
        AM: true
      },
      env_staging: {
        GO_ENV: "staging",
      },
      env_staging_am: {
        GO_ENV: "staging",
        AM: true
      },
      env_test: {
        GO_ENV: "test",
      },
      env_production: {
        GO_ENV: "production",
        AM: true
      }
    }]
  }