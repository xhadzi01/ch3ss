
function create_filled_piece_text(image_name, dev_id_str, figure_id_str, isDraggable){
  var filled_piece_template = `<div class="chessPieceDiv" id="DEV_ID_STR" DIV_ON_DRAG_PLACEHOLDER>
                                  <img class="chessPiece" src="/static/IMAGE_NAME" FIGURE_ON_DRAG_PLACEHOLDER id="FIGURE_ID_STR">
                              </div>`

  filled_piece_template = filled_piece_template.replace('IMAGE_NAME', image_name).replace('DEV_ID_STR', dev_id_str).replace('FIGURE_ID_STR', figure_id_str);

  if (isDraggable)  {
    filled_piece_template = filled_piece_template.replace('DIV_ON_DRAG_PLACEHOLDER', 'ondrop="drop(event)" ondragover="allowDrop(event)"').replace('FIGURE_ON_DRAG_PLACEHOLDER', 'draggable="true" ondragstart="drag(event)"');
  } else {
    filled_piece_template = filled_piece_template.replace('DIV_ON_DRAG_PLACEHOLDER', '').replace('FIGURE_ON_DRAG_PLACEHOLDER', '');
  }

  return filled_piece_template
}

function create_empty_piece_text(dev_id_str, isDraggable){
  var empty_piece_template = `<div class="chessPieceDiv" id="DEV_ID_STR" DIV_ON_DRAG_PLACEHOLDER></div>`

  empty_piece_template = empty_piece_template.replace('DEV_ID_STR', dev_id_str);

  if (isDraggable)  {
    empty_piece_template = empty_piece_template.replace('DIV_ON_DRAG_PLACEHOLDER', 'ondrop="drop(event)" ondragover="allowDrop(event)"');
  } else {
    empty_piece_template = empty_piece_template.replace('DIV_ON_DRAG_PLACEHOLDER', '');
  }
  
  return empty_piece_template
}


function refreshBoard(){

  var boardElement = document.getElementById('chess_piece_placement')

  boardElement.innerHTML = create_filled_piece_text("pawn_r.svg", "some_id1", "some_fig") + 
  create_filled_piece_text("pawn_r.svg", "some_id2", "some_id2_fig", true) +
  create_filled_piece_text("pawn_r.svg", "some_id3", "some_id3_fig", true) +
  create_filled_piece_text("pawn_r.svg", "some_id4", "some_id4_fig", true) +
  create_filled_piece_text("pawn_r.svg", "some_id5", "some_id5_fig", true) +
  create_filled_piece_text("pawn_r.svg", "some_id6", "some_id6_fig", true) +
  create_filled_piece_text("pawn_r.svg", "some_id7", "some_id7_fig", true) +
  create_filled_piece_text("pawn_r.svg", "some_id8", "some_id8_fig", true) +
  create_filled_piece_text("pawn_r.svg", "some_id9", "some_id9_fig", true) +
  create_filled_piece_text("pawn_r.svg", "some_id10", "some_id10_fig", true) +
  create_filled_piece_text("pawn_r.svg", "some_id11", "some_id11_fig", true) +
  create_filled_piece_text("pawn_r.svg", "some_id12", "some_id12_fig", true) +
  create_filled_piece_text("pawn_r.svg", "some_id13", "some_id13_fig", true) +
  create_filled_piece_text("pawn_r.svg", "some_id14", "some_id14_fig", true) +
  create_filled_piece_text("pawn_r.svg", "some_id15", "some_id15_fig", true) +
  create_filled_piece_text("pawn_r.svg", "some_id16", "some_id16_fig", true) +
  create_empty_piece_text("some_id17", true) +
  create_empty_piece_text("some_id18", true) +
  create_empty_piece_text("some_id19", true) +
  create_empty_piece_text("some_id20", true) +
  create_empty_piece_text("some_id21", true) +
  create_empty_piece_text("some_id22", true) +
  create_empty_piece_text("some_id23", true) +
  create_empty_piece_text("some_id24", true) +
  create_empty_piece_text("some_id25", true) +
  create_empty_piece_text("some_id26", true) +
  create_empty_piece_text("some_id27", true) +
  create_empty_piece_text("some_id28", true) +
  create_empty_piece_text("some_id29", true) +
  create_empty_piece_text("some_id30", false) +
  create_empty_piece_text("some_id31", false) +
  create_empty_piece_text("some_id32", false) +
  create_empty_piece_text("some_id33", false) +
  create_empty_piece_text("some_id34", false) +
  create_empty_piece_text("some_id35", false) +
  create_empty_piece_text("some_id36", false) +
  create_empty_piece_text("some_id37", false) +
  create_empty_piece_text("some_id38", false) +
  create_empty_piece_text("some_id39", false) +
  create_empty_piece_text("some_id40", false) +
  create_empty_piece_text("some_id41", false) +
  create_empty_piece_text("some_id42", false) +
  create_empty_piece_text("some_id43", false) +
  create_empty_piece_text("some_id44", false) +
  create_empty_piece_text("some_id45", false) +
  create_empty_piece_text("some_id46", false) +
  create_empty_piece_text("some_id47", false) +
  create_empty_piece_text("some_id48", false) +
  create_filled_piece_text("pawn_g.svg", "some_id49", "some_id49_fig", false) +
  create_filled_piece_text("pawn_g.svg", "some_id50", "some_id50_fig", false) +
  create_filled_piece_text("pawn_g.svg", "some_id51", "some_id51_fig", false) +
  create_filled_piece_text("pawn_g.svg", "some_id52", "some_id52_fig", false) +
  create_filled_piece_text("pawn_g.svg", "some_id53", "some_id53_fig", false) +
  create_filled_piece_text("pawn_g.svg", "some_id54", "some_id54_fig", false) +
  create_filled_piece_text("pawn_g.svg", "some_id55", "some_id55_fig", false) +
  create_filled_piece_text("pawn_g.svg", "some_id56", "some_id56_fig", false) +
  create_filled_piece_text("pawn_g.svg", "some_id57", "some_id57_fig", false) +
  create_filled_piece_text("pawn_g.svg", "some_id58", "some_id58_fig", false) +
  create_filled_piece_text("pawn_g.svg", "some_id59", "some_id59_fig", false) +
  create_filled_piece_text("pawn_g.svg", "some_id60", "some_id60_fig", false) +
  create_filled_piece_text("pawn_g.svg", "some_id61", "some_id61_fig", false) +
  create_filled_piece_text("pawn_g.svg", "some_id62", "some_id62_fig", false) +
  create_filled_piece_text("pawn_g.svg", "some_id63", "some_id63_fig", false) +
  create_filled_piece_text("pawn_g.svg", "some_id64", "some_id64_fig", false)
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

refreshBoard();



/* dragability events*/
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
