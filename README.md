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

```bash
make deploy
```

`aws cloudformation deploy`コマンドの`--parameter-overrides`オプションでパラメータの指定を行なっています。

このサンプルでは複数のパラメータを指定するためiniファイルに定義しました。

`parameters.sample.ini`を参考にしてください。
