# myhead

指定したファイルを再帰的に調査し、画像を変換します。


## Installation

`$ go get github.com/srttk/imgconv`

## Usage

`$ imgconv -src=png -out=jpg YOUR_WANT_CONVERT_IMAGE_DIR_NAME`

### src引数

`src` 属性は `png, jpg` などの変換したい画像の拡張子です。

この属性に指定したファイルタイプの画像のみ変換します。

対応している拡張子は `png, jpg, jpeg, gif` です。

### out属性

`out` 属性もおなじく、ファイルタイプを指定する拡張子を入力します。

この属性で指定した拡張子に画像を変換します。

対応している拡張子は `png, jpg, jpeg, gif` です。

## Author

Yuya Okumur (srttk)
