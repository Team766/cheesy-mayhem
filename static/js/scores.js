var patchScore = function (score) {
    fetch('/api/scores', {
        method: 'PATCH',
        body: JSON.stringify(score),
    });
    // TODO: add error handling
}

var toggleBoolean = function (id) {
    $(id).attr("data-value", $(id).attr("data-value") != "true");
}

var toggleTriState = function (id) {
    $(id).attr("data-value", ($(id).attr("data-value") + 1) % 3);
}

var decrementCount = function (id) {
    var curValue = parseInt($(id).text());
    $(id).text(Math.max(curValue - 1, 0));
}

var incrementCount = function (id) {
    var curValue = parseInt($(id).text());
    $(id).text(curValue + 1);
}

var parseBoolean = function (text) {
    return text != undefined && text === "true";
}

var getTaxi = function (alliance) {
    var taxi1 = parseInt($("#" + alliance + "Taxi1").attr("data-value"));
    var taxi2 = parseInt($("#" + alliance + "Taxi2").attr("data-value"));

    var score = {};
    score.taxi = [taxi1, taxi2];
    return { [alliance]: score };
}

var getShelf = function (alliance) {
    var autonBottom = parseInt($("#" + alliance + "AutonBottom").text());
    var autonTop = parseInt($("#" + alliance + "AutonTop").text());
    var teleopBottom = parseInt($("#" + alliance + "TeleopBottom").text());
    var teleopTop = parseInt($("#" + alliance + "TeleopTop").text());

    var score = {};
    score.shelf = {
        "auton_bottom": autonBottom,
        "auton_top": autonTop,
        "teleop_bottom": teleopBottom,
        "teleop_top": teleopTop,
    }
    return { [alliance]: score };
}

var getGoldenCube = function (alliance) {
    var goldenCube = parseBoolean($("#" + alliance + "GoldenCube").attr("data-value"));

    var score = {};
    score.golden_cube = goldenCube;
    return { [alliance]: score };
}

var getHamper = function (alliance) {
    var hamper = parseInt($("#" + alliance + "Hamper").text());
    var score = {};
    score.hamper = hamper;
    return { [alliance]: score };
}

var getPark = function (alliance) {
    var park1 = parseBoolean($("#" + alliance + "Park1").attr("data-value"));
    var park2 = parseBoolean($("#" + alliance + "Park2").attr("data-value"));

    var score = {};
    score.park = [park1, park2];
    return { [alliance]: score };
}

var getFoul = function (alliance) {
    var foul = parseInt($("#" + alliance + "Foul").text());
    var score = {};
    score.foul = foul;
    return { [alliance]: score };
}

var getTechFoul = function (alliance) {
    var techFoul = parseInt($("#" + alliance + "TechFoul").text());
    var score = {};
    score.tech_foul = techFoul;
    return { [alliance]: score };
}
