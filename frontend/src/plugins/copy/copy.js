import ClipboardJS from "clipboard";

/**
 * copy text
 * @param divClass
 */
export function copyText(divClass) {
  let boxs = document.querySelectorAll(divClass);
  if (boxs.length === 0) {
    return;
  }

  let normalIcon = "<span class='arco-btn-icon'><svg viewBox='0 0 48 48' fill='none' xmlns='http://www.w3.org/2000/svg' stroke='currentColor' class='arco-icon arco-icon-copy' stroke-width='4' stroke-linecap='butt' stroke-linejoin='miter' filter=''><path d='M20 6h18a2 2 0 0 1 2 2v22M8 16v24c0 1.105.891 2 1.996 2h20.007A1.99 1.99 0 0 0 32 40.008V15.997A1.997 1.997 0 0 0 30 14H10a2 2 0 0 0-2 2Z'></path></svg></span>";
  let copiedIcon = "<span class='arco-btn-icon'><svg viewBox='0 0 48 48' fill='none' xmlns='http://www.w3.org/2000/svg' stroke='currentColor' class='arco-icon arco-icon-check' stroke-width='4' stroke-linecap='butt' stroke-linejoin='miter' filter='' style='color:#21b42a;'><path d='M41.678 11.05 19.05 33.678 6.322 20.95'></path></svg></span>";
  // create btn
  let copyBtn = document.createElement("button");
  copyBtn.className = "arco-btn arco-btn-secondary hljs-btn-copy";
  copyBtn.innerHTML = normalIcon;
  copyBtn.style.position = 'absolute';
  copyBtn.style.top = '10px';
  copyBtn.style.right = '10px';

  boxs.forEach((box, index) => {
    // box event
    box.addEventListener("mouseover", function (e) {
      copyBtn.className += " hljs-btn" + index
      copyBtn.setAttribute("data-clipboard-text", box.innerText);
      copyBtn.setAttribute('data-clipboard-action', 'copy')
      box.appendChild(copyBtn);
      copyBtn.style.display = 'block';
      copyBtn.style.pointerEvents = 'auto';

      // button event
      copyBtn.addEventListener('mouseover', function (e) {
        copyBtn.style.pointerEvents = 'auto';
      });
      copyBtn.addEventListener('mouseout', function (e) {
        copyBtn.style.pointerEvents = 'none';
      });
      copyBtn.addEventListener('click', function (e) {
        const clipboard = new ClipboardJS('.hljs-btn' + index);
        clipboard.on('success', function (e) {
          copyBtn.innerHTML = copiedIcon;
          setTimeout(() => {
            copyBtn.innerHTML = normalIcon;
          }, 1500);
        });
        clipboard.on('error', function (e) {
          console.info('Action:', e.action);
          console.info('Text:', e.text);
          console.info('Trigger:', e.trigger);
        });
      });
    });
    box.addEventListener('mouseout', function (e) {
      copyBtn.style.display = 'none';
    });
  })
}

