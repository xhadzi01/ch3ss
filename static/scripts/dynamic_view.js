
function create_filled_piece_text(image_name, id_str){
    var filled_piece_template = `<div id="div1" class="chessPieceDiv" id="ID_STR" ondrop="drop(event)" ondragover="allowDrop(event)">
                                    <img class="chessPiece" src="/static/IMAGE_NAME" draggable="true" ondragstart="drag(event)" id="drag1">
                                </div>`

    return filled_piece_template.replace('IMAGE_NAME', image_name).replace('ID_STR', id_str);
}

function create_empty_piece_text(id_str){
    var empty_piece_template = `<div class="chessPieceDiv" id="ID_STR" ondrop="drop(event)" ondragover="allowDrop(event)"></div>`
    return empty_piece_template.replace('ID_STR', id_str);
}


function refreshBoard(){

    var boardElement = document.getElementById('chess_piece_placement')

    boardElement.innerHTML = create_filled_piece_text("pawn_r.svg", "some_id") + 
    create_filled_piece_text("pawn_r.svg", "some_id") +
    create_filled_piece_text("pawn_r.svg", "some_id") +
    create_filled_piece_text("pawn_r.svg", "some_id") +
    create_filled_piece_text("pawn_r.svg", "some_id") +
    create_filled_piece_text("pawn_r.svg", "some_id") +
    create_filled_piece_text("pawn_r.svg", "some_id") +
    create_filled_piece_text("pawn_r.svg", "some_id") +
    create_filled_piece_text("pawn_r.svg", "some_id") +
    create_filled_piece_text("pawn_r.svg", "some_id") +
    create_filled_piece_text("pawn_r.svg", "some_id") +
    create_filled_piece_text("pawn_r.svg", "some_id") +
    create_filled_piece_text("pawn_r.svg", "some_id") +
    create_filled_piece_text("pawn_r.svg", "some_id") +
    create_filled_piece_text("pawn_r.svg", "some_id") +
    create_filled_piece_text("pawn_r.svg", "some_id") +
    create_empty_piece_text("some_id") +
    create_empty_piece_text("some_id") +
    create_empty_piece_text("some_id") +
    create_empty_piece_text("some_id") +
    create_empty_piece_text("some_id") +
    create_empty_piece_text("some_id") +
    create_empty_piece_text("some_id") +
    create_empty_piece_text("some_id") +
    create_empty_piece_text("some_id") +
    create_empty_piece_text("some_id") +
    create_empty_piece_text("some_id") +
    create_empty_piece_text("some_id") +
    create_empty_piece_text("some_id") +
    create_empty_piece_text("some_id") +
    create_empty_piece_text("some_id") +
    create_empty_piece_text("some_id") +
    create_empty_piece_text("some_id") +
    create_empty_piece_text("some_id") +
    create_empty_piece_text("some_id") +
    create_empty_piece_text("some_id") +
    create_empty_piece_text("some_id") +
    create_empty_piece_text("some_id") +
    create_empty_piece_text("some_id") +
    create_empty_piece_text("some_id") +
    create_empty_piece_text("some_id") +
    create_empty_piece_text("some_id") +
    create_empty_piece_text("some_id") +
    create_empty_piece_text("some_id") +
    create_empty_piece_text("some_id") +
    create_empty_piece_text("some_id") +
    create_empty_piece_text("some_id") +
    create_empty_piece_text("some_id") +
    create_filled_piece_text("pawn_g.svg", "some_id") +
    create_filled_piece_text("pawn_g.svg", "some_id") +
    create_filled_piece_text("pawn_g.svg", "some_id") +
    create_filled_piece_text("pawn_g.svg", "some_id") +
    create_filled_piece_text("pawn_g.svg", "some_id") +
    create_filled_piece_text("pawn_g.svg", "some_id") +
    create_filled_piece_text("pawn_g.svg", "some_id") +
    create_filled_piece_text("pawn_g.svg", "some_id") +
    create_filled_piece_text("pawn_g.svg", "some_id") +
    create_filled_piece_text("pawn_g.svg", "some_id") +
    create_filled_piece_text("pawn_g.svg", "some_id") +
    create_filled_piece_text("pawn_g.svg", "some_id") +
    create_filled_piece_text("pawn_g.svg", "some_id") +
    create_filled_piece_text("pawn_g.svg", "some_id") +
    create_filled_piece_text("pawn_g.svg", "some_id") +
    create_filled_piece_text("pawn_g.svg", "some_id")
}

function setChessPieceLocation(location, pieceName){
    var boardElement = document.getElementById('chess_piece_placement')
}

document.getElementById('btn').onclick = function() {
    //var val = document.getElementById('imagename').value,
    //    src = '/static/' + val +'.svg',
    //    img = document.createElement('img');

    //img.src = src;
    //document.body.appendChild(img);

    refreshBoard()
    //document.getElementById('chess_piece_placement').appendChild(img);

}


function allowDrop(ev) {
  ev.preventDefault();
}

function drag(ev) {
  ev.dataTransfer.setData("text", ev.target.id);
}

function drop(ev) {
  ev.preventDefault();
  var data = ev.dataTransfer.getData("text");
  ev.target.appendChild(document.getElementById(data));
}
