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
    // update the overall score
    if (alliance === "red") {
        $("#score").text(data.Red.ScoreSummary.Score);
    } else if (alliance === "blue") {
        $("#matchScore").text(data.Blue.ScoreSummary.Score);
    }

    // 
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
