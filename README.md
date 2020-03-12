# aws-stepfunctions-samples

AWS StepFunctionsを利用してFargateのタスクを起動するサンプルです。


CloudFormationによるデプロイを行います。

## 前提条件

- AWS CLI v1


## デプロイ

### ECRへのpush

AWSアカウントIDを第1引数に指定して下さい
```bash
push-ecr.sh 123456789012
```

### CloudFormationによるデプロイ

makeコマンドでデプロイを行います。

必要に応じて`--parameters`でオプションを指定してください。

```bash
make deploy
```
