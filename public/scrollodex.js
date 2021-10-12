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

function ellipsisRenderer(str, cellRef, $cell) {
  num = 300;
  if (str.length > num) {
    return str.slice(0, num) + "...";
  }
  return str;
}
ZingGrid.registerCellType("ellipsis", { renderer: ellipsisRenderer });
