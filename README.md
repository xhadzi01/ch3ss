# ch3ss
Simple chess written in go using basic Http for user interaction.

### Expected user interactions:
1.) Player1 presses "Start a New Game", he is stuck on waiting screen until Player2 joins. He will be displayed Session Code, which is used to join the game.
2.) Player2 writes Session Code into a text box and clicks on "Join Game".
3.) Both players are shown chess board with the default chess piece location. Black and White players are randomly selected.
4.) White player starts. Mouse clicks on a piece and drags it onto other position that is highlighted as available. Black player waits on his turn. 
5.) Black player continues. White waits on his turn.
6.) Players continue with alternation until one wins the game.
7.) Game ends. Basic score board is shown.

### Technical details:
ad. 1.) On clicking on "Start a New Game", a player is assiged a Player ID which will be linked to Session ID. Player is shown a Session Code, whitch is also linked to Session ID. This code could be used to connect to the same game. Second player will use te Session code to connect to the game session. 
ad. 2.) When the seconds player fills in a Session code and clicks on "Join Game". Game engine check whether the Session is valid:
- Session is still active
- Session has only one player
If the session is valid, he will be assigned a Player ID. Player ID will be linked to the Session ID.
ad. 3.) -
ad. 4.) Each piece will have its permitted moves. Grabbing it with the mouse will highlight available positions. Validation of the position will be also done on pieces on mouse release.
