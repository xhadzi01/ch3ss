document.getElementById('btn').onclick = function() {
    var val = document.getElementById('imagename').value,
        src = '/static/' + val +'.svg',
        img = document.createElement('img');

    img.src = src;
    document.body.appendChild(img);
}