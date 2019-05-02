# health-to-csv

## Overview

iOSのヘルスデータをCSVに変換するツール

## Description

iOSのヘルスデータ（xml）をCSV形式に変換します。
データ量が多いとそれだけメモリを食うので注意です。

CSVデータに出力する項目は、
export.xmlのRecordデータから、以下のみに制限しています。
* type
* unit
* startDate/endDate
* value

## Usage

* iPhoneのヘルスアプリから、データをexportする
* カレントディレクトリにzipファイルを解凍する
  ```bash
    .
    ├── README.md
    ├── apple_health_export
    │   ├── export.xml
    │   └── export_cda.xml
  ```
* アプリケーションを実行する

