/*!
 * Scrollodex JavaScript
 */


function upperRenderer(mRawData, cellRef, $cell) {
  return mRawData.toUpperCase();
}
ZingGrid.registerCellType('upper', { renderer: upperRenderer, });


function displayLocRenderer(CountryCode, Region, Comment, cellRef, $cell) {
  if (Comment != "") {
    return CountryCode + "-" + Region + " (" + Comment + ")";
  } else {
    return CountryCode + "-" + Region;
  }
}
ZingGrid.registerCellType('displayLoc', { renderer: displayLocRenderer, });
ZingGrid.registerCellType('upper', { renderer: upperRenderer, });
