// Copyright 2014 Team 254. All Rights Reserved.
// Author: pat@patfairbank.com (Patrick Fairbank)
//
// Client-side methods for editing a match in the match review page.

var scoreTemplate = Handlebars.compile($("#scoreTemplate").html());
var allianceResults = {};
var matchResult;

// Hijack the form submission to inject the data in JSON form so that it's easier for the server to parse.
$("form").submit(function() {
  updateResults("red");
  updateResults("blue");

  matchResult.RedScore = allianceResults["red"].score;
  matchResult.BlueScore = allianceResults["blue"].score;
  var matchResultJson = JSON.stringify(matchResult);

  // Inject the JSON data into the form as hidden inputs.
  $("<input />").attr("type", "hidden").attr("name", "matchResultJson").attr("value", matchResultJson).appendTo("form");

  return true;
});

// Draws the match-editing form for one alliance based on the cached result data.
var renderResults = function(alliance) {
  var result = allianceResults[alliance];
  var scoreContent = scoreTemplate(result);
  $("#" + alliance + "Score").html(scoreContent);

  getInputElement(alliance, "Taxi1").val(result.score.Taxi[0]);
  getInputElement(alliance, "Taxi2").val(result.score.Taxi[1]);
  getInputElement(alliance, "AutonBottom").val(result.score.Shelf.AutonBottomShelfCubes);
  getInputElement(alliance, "AutonTop").val(result.score.Shelf.AutonTopShelfCubes);
  getInputElement(alliance, "TeleopBottom").val(result.score.Shelf.TeleopBottomShelfCubes);
  getInputElement(alliance, "TeleopTop").val(result.score.Shelf.TeleopTopShelfCubes);
  getInputElement(alliance, "GoldenCube").prop('checked', result.score.GoldenCube);
  getInputElement(alliance, "Hamper").val(result.score.Hamper);
  getInputElement(alliance, "Park1").prop('checked', result.score.Park[0]);
  getInputElement(alliance, "Park2").prop('checked', result.score.Park[1]);
};

// Converts the current form values back into JSON structures and caches them.
var updateResults = function(alliance) {
  var result = allianceResults[alliance];
  var formData = {};
  $.each($("form").serializeArray(), function(k, v) {
    formData[v.name] = v.value;
  });

  result.score.Taxi = [ parseInt(formData[alliance + "Taxi1"]), parseInt(formData[alliance + "Taxi2"]) ];
  result.score.Shelf.AutonBottomShelfCubes = parseInt(formData[alliance + "AutonBottom"]);
  result.score.Shelf.AutonTopShelfCubes = parseInt(formData[alliance + "AutonTop"]);
  result.score.Shelf.TeleopBottomShelfCubes = parseInt(formData[alliance + "TeleopBottom"]);
  result.score.Shelf.TeleopTopShelfCubes = parseInt(formData[alliance + "TeleopTop"]);
  result.score.GoldenCube = parseBool(formData[alliance + "GoldenCube"]);
  result.score.Hamper = parseInt(formData[alliance + "Hamper"]);
  result.score.Park = [ parseBool(formData[alliance + "Park1"]), parseBool(formData[alliance + "Park2"]) ];
};

// Returns the form input element having the given parameters.
var getInputElement = function(alliance, name, value) {
  var selector = "input[name=" + alliance + name + "]";
  if (value !== undefined) {
    selector += "[value=" + value + "]";
  }
  return $(selector);
};

// TODO: is there a builtin way to do this?
var parseBool = function(text) {
  if (text === undefined) {
    return false;
  }
  var lower = text.toLowerCase();
  var num = parseInt(text);
  if (lower === "true" || num == 1) {
    return true;
  } else if (lower === "false" || num == 0) {
    return false;
  }
  return false;
};
