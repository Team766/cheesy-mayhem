// Client-side logic for the ref page.

var websocket;
let alliance;

var updateScore = function(score) {
    // update each of the individal scoring elements
    // taxi
    $("#" + alliance + "Taxi1").attr("data-value", score.Taxi[0]);
    $("#" + alliance + "Taxi2").attr("data-value", score.Taxi[1]);

    // shelf
    $("#" + alliance + "AutonBottom").text(score.Shelf.AutonBottomShelfCubes);
    $("#" + alliance + "AutonTop").text(score.Shelf.AutonTopShelfCubes);
    $("#" + alliance + "TeleopBottom").text(score.Shelf.TeleopBottomShelfCubes);
    $("#" + alliance + "TeleopTop").text(score.Shelf.TeleopTopShelfCubes);

    // golden_cube
    $("#" + alliance + "GoldenCube").attr("data-value", score.GoldenCube);

    // hamper
    $("#" + alliance + "Hamper").text(score.Hamper);

    // park
    $("#" + alliance + "Park1").attr("data-value", score.Park[0]);
    $("#" + alliance + "Park2").attr("data-value", score.Park[1]);

    // foul
    $("#" + alliance + "Foul").text(score.Fouls);

    // tech foul
    $("#" + alliance + "TechFoul").text(score.TechFouls);
}

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
        $("#matchScore").text(data.Red.ScoreSummary.Score);
        score = data.Red.Score;
    } else if (alliance === "blue") {
        $("#matchScore").text(data.Blue.ScoreSummary.Score);
        score = data.Blue.Score;
    }
    updateScore(score);
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
