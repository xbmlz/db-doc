package doc

const mdCss = `<style type="text/css">
body {
	width: 210mm;
	margin: auto;
}
table {
	width: 100%;
	border-collapse: collapse;
	border-spacing: 0;
}
table tr {
	background-color: #fff;
	border-top: 1px solid #ccc;
}
table td ,table th {
	padding: 6px 13px;
	border: 1px solid #ddd;
}
table th{
	background-color: rgb(64, 158, 255);
	color: rgb(255, 255, 255);
}
</style>
`

const docsifyHTML = `<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<title>Database Document</title>
<meta http-equiv="X-UA-Compatible" content="IE=edge,chrome=1" />
<meta name="description" content="Description">
<meta name="viewport" content="width=device-width, user-scalable=no, initial-scale=1.0, maximum-scale=1.0, minimum-scale=1.0">
<link rel="stylesheet" href="//cdn.staticfile.org/docsify/4.12.1/themes/vue.min.css">
</head>
<body>
<div data-app id="main">加载中</div>
<script>
	window.$docsify = {
		el: '#main',
		name: '',
		repo: '',
		search: 'auto',
		loadSidebar: true
	}
</script>
<script src="//cdn.staticfile.org/docsify/4.12.1/docsify.min.js"></script>
<script src="//cdn.staticfile.org/docsify/4.12.1/plugins/search.min.js"></script>
</body>
</html>
`
