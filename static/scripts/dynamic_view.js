/////////////////////////////////////////////////////
// fetch requests

function get_cookie(name) {
  const value = `; ${document.cookie}`;
  const parts = value.split(`; ${name}=`);
  if (parts.length === 2) return parts.pop().split(';').shift();
}

function redraw_current_figure_positions() {
  console.debug('fetching game figure positions')
  let session_id_str = get_cookie("sessionID")
  fetch(`/get-game-info/${session_id_str}`, {
      credentials: "same-origin",
      mode: "same-origin",
      method: "GET",
      headers: { "Content-Type": "application/json" }
  }).then(resp => {
      if (resp.status === 200) {
          return resp.json();
      } else {
          throw new Error('Something went wrong');
      }
  }).then(dataJson => {
    console.debug('fetching game figure positions - success')
    draw_chessboard_figures(dataJson.player1.FigureInfoList, dataJson.player2.FigureInfoList)
  }).catch(err => {
      console.log(err);
  })
}

function send_figure_position_update(figure, position) {
  console.debug(`setting figure: '${figure}' to position: '${position}`)
  let session_id_str = get_cookie("sessionID")
  let success = false
  fetch(`/move_figure/${session_id_str}`, {
      credentials: "same-origin",
      mode: "same-origin",
      method: "PUT",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({figure: figure, targetPosition: position})
  }).then(resp => {
      if (resp.status === 200) {
          return resp.json();
      } else {
          throw new Error('Something went wrong');
      }
  }).then(dataJson => {
    console.debug('moving figure was sucessfull')
    success = true
  }).catch(err => {
      console.log(err);
  })

  return success
}

/////////////////////////////////////////////////////
// figure image selection

const figure_type_pawn   = 1
const figure_type_rook   = 2
const figure_type_knight = 3
const figure_type_bishop = 4
const figure_type_queen  = 5
const figure_type_king   = 6

function get_player_figure_img(figure_type, player_name) {
  return `${get_player_figure_type_str(figure_type)}_${get_player_figure_color(player_name)}.svg`
}

function get_player_figure_type_str(figure_type) {
  switch (figure_type) {
    case figure_type_pawn:
      return `pawn`
    case figure_type_rook:
      return `rook`
    case figure_type_knight:
      return `knight`
    case figure_type_bishop:
      return `bishop`
    case figure_type_queen:
      return `queen`
    case figure_type_king:
      return `king`
  }
}

function get_player_figure_color(player_name) {
  if (player_name === `player1`) {
    return `g`
  } else {
    return `r`
  }
}

/////////////////////////////////////////////////////
// calculation of full chessboard

function create_chessboard_position_map(player1_figures, player2_figures) {
  let chessboard_positions = {}
  // there is 64 chess positions (8x8 grid)
  for (let i = 1; i <= 64; i++) {
    chessboard_positions[i] = null
  }

  // then iterate over player 1 and add figures to position map
  for (let i = 0; i < player1_figures.length; i++) {
    player_figure = player1_figures[i]
    chessboard_positions[player_figure.CurrentFigurePosition] = {
      figure_id: `figure_player1_${player_figure.FigureIndex}_id`,
      figure_img_src: get_player_figure_img(player_figure.FigureType, 'player1'),
    }
  }

  // then iterate over player 2 and add figures to position map
  for (let i = 0; i < player2_figures.length; i++) {
    player_figure = player2_figures[i]
    chessboard_positions[player_figure.CurrentFigurePosition] = {
      figure_id: `figure_player2_${player_figure.FigureIndex}_id`,
      figure_img_src: get_player_figure_img(player_figure.FigureType, 'player2'),
    }
  }
  
  return chessboard_positions
}

/////////////////////////////////////////////////////
// element drawing

function create_filled_piece_text(image_name, dev_id_str, figure_id_str, index, isDraggable){
  console.debug('create_filled_piece_text - enter')

  let dragable_event_text = ''
  let figure_on_drag_event_text = ''
  if (isDraggable)  {
    dragable_event_text = `ondrop="drop(event, ${dev_id_str})" ondragover="allowDrop(event, ${dev_id_str})"`
    figure_on_drag_event_text = `draggable="true" ondragstart="drag(event, ${dev_id_str})"`
  }

  var filled_piece_template = `<div class="chessPieceDiv" id="${dev_id_str}" ${dragable_event_text}>
                                  <img class="chessPiece" src="/static/${image_name}" ${figure_on_drag_event_text} id="${figure_id_str}">
                              </div>`

  return filled_piece_template
}

function create_empty_piece_text(dev_id_str, index, isDraggable){
  console.debug('create_empty_piece_text - enter')
  
  let dragable_event_text = ''
  if (isDraggable)  {
    dragable_event_text = `ondrop="drop(event, ${dev_id_str})" ondragover="allowDrop(event)"`
  }
  
  var empty_piece_template = `<div class="chessPieceDiv" id="${dev_id_str}" ${dragable_event_text}></div>`

  return empty_piece_template
}

function draw_chessboard_figures(player1_figures, player2_figures) {
  console.debug('drawing chestboard figures - enter')
  let position_map = create_chessboard_position_map(player1_figures, player2_figures);

  var boardElement = document.getElementById('chess_piece_placement')
  let html_element_str = ""
  for (let i = 1; i <= 64; i++) {
    if (position_map[i] != null) {
      html_element_str += create_filled_piece_text(position_map[i].figure_img_src, `place_${i}_id`, position_map[i].figure_id, i, true);
    } else {
      html_element_str += create_empty_piece_text(`place_${i}_id`, i, true);
    }
  }

  //console.debug('drawing chestboard figures - raw text:', html_element_str)
  boardElement.innerHTML = html_element_str
}

/////////////////////////////////////////////////////
// events

// is executed when figure is grabbed
function drag(ev, dragged_figure) {
  ev.dataTransfer.setData("text", ev.target.id);
  console.debug(`dragged event - place_id: '${dragged_figure.id}' figure_id: '${ev.target.id}'`)
}

// is executed when figure is hovered over a field
function allowDrop(ev, target_slot) {
  ev.preventDefault();
  console.debug(`drag-over - target_place_id: '${target_slot}'`)
}

// is executed when figure is dropped
function drop(ev, target_slot) {
  ev.preventDefault();
  var dropped_figure_id = ev.dataTransfer.getData("text");
  console.debug(`dropped - target_place_id: '${target_slot.id}' dropped_figure_id: '${dropped_figure_id}'`)
  send_figure_position_update(dropped_figure_id, target_slot.id)

  ev.target.appendChild(document.getElementById(dropped_figure_id));
}


/////////////////////////////////////////////////////
// on startup
redraw_current_figure_positions()
