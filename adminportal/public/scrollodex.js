/*!
 * Scrollodex JavaScript
 */

function displayLocRenderer(CountryCode, Region, Comment, cellRef, $cell) {
  // This must match dextidy.MakeDisplayLoc. If you change this,
  // change it too.

  if (CountryCode === "ZZ") {
    if (Comment === "") {
      return Region;
    } else {
      return Region + " (" + Comment + ")";
    }
  }
  if (Comment === "") {
    return CountryCode + "-" + Region;
  }
  return CountryCode + "-" + Region + " (" + Comment + ")";
}
ZingGrid.registerCellType("displayLoc", { renderer: displayLocRenderer });
