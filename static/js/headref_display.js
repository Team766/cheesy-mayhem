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
    $("#" + alliance + "Foul").text(score.Foul);

    // tech foul
    $("#" + alliance + "TechFoul").text(score.TechFoul);
}

// Handles a websocket message to update the match score.
var handleRealtimeScore = function(data) {
    $("#redScore").text(data.Red.ScoreSummary.Score);
    $("#blueScore").text(data.Blue.ScoreSummary.Score);

    updateAllianceScore("red", data.Red.Score);
    updateAllianceScore("blue", data.Blue.Score);
};

$(function() {
    // Set up the websocket back to the server.
    websocket = new CheesyWebsocket("/displays/headref/websocket", {
        matchLoad: function (event) { handleMatchLoad(event.data); },
        matchTiming: function (event) { handleMatchTiming(event.data); },
        matchTime: function (event) { handleMatchTime(event.data); },
        realtimeScore: function (event) { handleRealtimeScore(event.data); },
    });
});
