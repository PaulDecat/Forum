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