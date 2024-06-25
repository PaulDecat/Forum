fetch('createp.html')
  .then(response => response.text())
  .then(data => {
    
    const title = data.match(/<h1>(.*?)<\/h1>/)[1];
    const text = data.match(/<div id="creatingmypost">(.*?)<\/div>/)[1];

    const tp = new XMLHttpRequest();
    tp.open('GET', 'index.html', true);
    tp.onload = function () {
      if (this.status == 200) {
        const totalPostDiv = document.querySelector('#totalPost');
        totalPostDiv.innerHTML = `
          <h3>${title}</h3>
          <div id="creatingmypost">${text}</div>
        `;
      }
    };
    tp.send();
  });

/*
  fetch('createp.html')
  .then(response => response.text())
  .then(data => {
    const likeCheckboxChecked = /<input type="checkbox" id="mylikes" name="mylikes"(.*?)>/.test(data) && data.match(/<input type="checkbox" id="mylikes" name="mylikes"(.*?)>/)[1].includes('checked');
    const dislikeCheckboxChecked = /<input type="checkbox" id="mydislikes" name="mydislikes"(.*?)>/.test(data) && data.match(/<input type="checkbox" id="mydislikes" name="mydislikes"(.*?)>/)[1].includes('checked');

    const tp = new XMLHttpRequest();
    tp.open('GET', 'index.html', true);
    tp.onload = function () {
      if (this.status == 200) {
        const likeCheckbox = document.querySelector('#mylikes');
        if (likeCheckbox) {
          likeCheckbox.checked = likeCheckboxChecked;
          
        }

        const dislikeCheckbox = document.querySelector('#mydislikes');
        if (dislikeCheckbox) {
          dislikeCheckbox.checked = dislikeCheckboxChecked;
        }
      }
    };
    tp.send();
  })
  .catch(error => console.error('Error:', error));
*/