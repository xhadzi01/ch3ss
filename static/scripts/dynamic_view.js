document.getElementById('btn').onclick = function() {
    var val = document.getElementById('imagename').value,
        src = 'http://localhost:32000/static/' + val +'.svg',
        img = document.createElement('img');

    img.src = src;
    document.body.appendChild(img);
}