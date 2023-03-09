/**
 * CSS to hide everything on the page,
 * except for elements that have the "beastify-image" class.
 */
const hidepage = `body > :not(.beastify-image) {
  display: none;
}`;

/**
 * Listen for clicks on the buttons, and send the appropriate message to
 * the content script in the page.
 */
function listenForClicks() {
  document.addEventListener("click", () => {
    /**
     * Given the name of a beast, get the URL to the corresponding image.
     */
    function beastNameToURL(beastName) {
      switch (beastName) {
        case "Frog":
          return browser.runtime.getURL("beasts/frog.jpg")
        case "Snake":
          return browser.runtime.getURL("beasts/snake.jpg")
        case "Turtle":
          return browser.runtime.getURL("beasts/turtle.jpg")
      }
    }

    /**
     * Insert the page-hiding CSS into the active tab,
     * then get the beast URL and
     * send a beastify message to the content script in the active tab.
     */
    function beastify(tabs) {
      browser.tabs.insertCSS({code: hidepage}).then(() => {
        let url = beastNameToURL(e.target.textContent);
        browser.tabs.sendMessage(tabs[0].id, {
          command: "beastify",
          beastURL: url
        });
        console.log(tabs[0])
        console.log(tabs[0])
      });
    }

    /**
     * Remove the page-hiding css from the active tab,
     * send a "reset" message to the content script in the active tab.
     */
    function reset(tabs) {
      browser.tabs.removeCSS({code: hidepage}).then(() => {
        browser.tabs.sendMessage(tabs[0].id, {
          command: "reset",
        });
      });
    }

    /**
     * Just log the error to console.
     */
    function reportError(error) {
      console.error(`Could not beastify: ${error}`)
    }

    /**
     * Get the active tab,
     * then call "beastify()" or "reset()" as appropriate.
     */
    if (e.target.type === "reset") {
      browser.tabs
        .query({active: true, currentWindow: true})
        .then(reset)
        .catch(reportError)
    } else {
      browser.tabs
        .query({active: true, currentWindow: true})
        .then(beastify)
        .catch(reportError)
    }
  });
}

/**
 * There was an error executing the script.
 * Display the popup's error message, and hide the normal UI.
 */
function reportExecuteScriptError(error) {
  document.querySelector("#popup-content").classList.add("hidden");
  document.querySelector("#error-content").classList.add("hidden");
  console.error(`Failed to execute beastify content script: ${error.message}`)
}

/**
 * When the popup loads, inject a content script into the active tab,
 * and add a click handler.
 * If we couldn't inject the script, handler the error
 */
browser.tabs
  .executeScript({file:"/content_scripts/beastify.js"})
  .then(listenForClicks)
  .catch(reportExecuteScriptError)