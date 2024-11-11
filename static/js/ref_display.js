// Client-side logic for the ref page.

var websocket;
let alliance;

// Handles a websocket message to update the teams for the current match.
var handleMatchLoad = function(data) {
    $("#matchName").text(data.MatchType + " Match " + data.Match.DisplayName);
}

// Handles a websocket message to update the match time countdown.
var handleMatchTime = function(data) {
    translateMatchTime(data, function (matchState, matchStateText, countdownSec) {
        $("#matchState").text(matchStateText);
        $("#matchTime").text(countdownSec);
    });
};

// Handles a websocket message to update the match score.
var handleRealtimeScore = function(data) {

    var score;

    // update the overall score
    if (alliance === "red") {
        $("#score").text(data.Red.ScoreSummary.Score);
        score = data.Red.Score;
    } else if (alliance === "blue") {
        $("#matchScore").text(data.Blue.ScoreSummary.Score);
        score = data.Blue.Score;
    }

    // update each of the individal scoring elements
    // taxi
    $("#taxi1").text(score.Taxi[0]);
    $("#taxi2").text(score.Taxi[1]);

    // shelf
    $("#auton_bottom").text(score.Shelf.AutonBottomShelfCubes);
    $("#auton_top").text(score.Shelf.AutonTopShelfCubes);
    $("#teleop_bottom").text(score.Shelf.TeleopBottomShelfCubes);
    $("#teleop_top").text(score.Shelf.TeleopTopShelfCubes);

    // golden_cube
    $("#golden_cube").text(score.GoldenCube);

    // hamper
    $("#hamper").text(score.Hamper);

    // park
    $("#park1").text(score.Park[0]);
    $("#park2").text(score.Park[1]);
};

$(function() {
    alliance = window.location.href.split("/").slice(-1)[0];
    // Set up the websocket back to the server.
    websocket = new CheesyWebsocket("/displays/ref/" + alliance + "/websocket", {
        matchLoad: function (event) { handleMatchLoad(event.data); },
        matchTiming: function (event) { handleMatchTiming(event.data); },
        matchTime: function (event) { handleMatchTime(event.data); },
        realtimeScore: function (event) { handleRealtimeScore(event.data); },
    });
});
