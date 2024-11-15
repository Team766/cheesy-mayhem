// Client-side logic for the ref page.

var websocket;

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

var updateAllianceScore = function(alliance, score) {
    // update each of the individal scoring elements
    // taxi
    $("#" + alliance + "Taxi1").text(score.Taxi[0]);
    $("#taxi2").text(score.Taxi[1]);

    // shelf
    $("#" + alliance + "AutonBottom").text(score.Shelf.AutonBottomShelfCubes);
    $("#" + alliance + "AutonTop").text(score.Shelf.AutonTopShelfCubes);
    $("#" + alliance + "TeleopBottom").text(score.Shelf.TeleopBottomShelfCubes);
    $("#" + alliance + "TeleopTop").text(score.Shelf.TeleopTopShelfCubes);

    // golden_cube
    $("#" + alliance + "GoldenCube").text(score.GoldenCube);

    // hamper
    $("#" + alliance + "Hamper").text(score.Hamper);

    // park
    $("#" + alliance + "Park1").text(score.Park[0]);
    $("#" + alliance + "Park2").text(score.Park[1]);
}

// Handles a websocket message to update the match score.
var handleRealtimeScore = function(data) {

    var score;

    $("#redScore").text(data.Red.ScoreSummary.Score);
    $("#blueScore").text(data.Blue.ScoreSummary.Score);

    updateScore(data.Red.ScoreSummary.Score, "red");
    updateScore(data.Blue.ScoreSummary.Score, "blue");
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
