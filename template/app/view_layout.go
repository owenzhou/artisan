package app

var ViewLayoutTemplate = `
<!DOCTYPE html>
<html>
<head>
    <meta charset="utf-8">
    <title>首页</title>
</head>
<body>
    <div>
      {{ $a := "{{ template \"content\" .}}" }}
	  {{- $a }}
    </div>
</body>
</html>
`