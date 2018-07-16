(function() {
  var table = document.getElementById('table');
  var addButton = document.getElementById('add');
  addButton.addEventListener('click', function(e) {
    var rowData = createRow();

    rowData.actionButton.addEventListener('click', function save() {
      rowData.actionButton.removeEventListener('click', save);

      var payload = {
        startDate: rowData.startDate.value,
        hours: parseFloat(rowData.hours.value) || 0,
        reason: rowData.reasons.value,
        status: 'new'
      };

      rowData.row.parentElement.removeChild(rowData.row);

    }, false);

    table.appendChild(rowData.row);
  }, false);

  function createRow() {
    var result = {};

    var row = document.createElement('div');
    row.classList.add('row', 'col-lg-12');
    result.row = row;

    var dateCell = document.createElement('div');
    dateCell.classList.add('col-lg-3');
    var input = document.createElement('input');
    input.classList.add('form-control');
    input.type = 'date';
    result.startDate = input;
    dateCell.appendChild(input);

    var hoursCell = document.createElement('div');
    hoursCell.classList.add('col-lg-2');
    var input = document.createElement('input');
    input.classList.add('form-control');
    input.type = 'number';
    result.hours = input;
    hoursCell.appendChild(input);

    var reasonCell = document.createElement('div');
    reasonCell.classList.add('col-lg-2');
    var select = document.createElement('select');
    select.classList.add('form-control');
    var optionData = [
      ['pto', 'PTO'],
      ['leave', 'Leave of Absence'],
      ['family', 'Family Leave'],
      ['bereavement', 'Bereavement']
    ];
    optionData.forEach(function(pair){
      var opt = document.createElement('option');
      opt.value = pair[0];
      opt.text = pair[1];
      select.add(opt, null);
    });
    result.reasons = select;
    reasonCell.appendChild(select);

    var statusCell = document.createElement('div');
    statusCell.classList.add('col-lg-2', 'text-center');
    statusCell.innerHTML = 'New';
    result.status= statusCell;

    var actionCell = document.createElement('div');
    actionCell.classList.add('col-lg-3');
    var btn = document.createElement('button');
    btn.classList.add('btn', 'primary');
    btn.innerHTML = 'Save';
    result.actionButton = btn;
    actionCell.appendChild(btn);

    row.appendChild(dateCell);
    row.appendChild(hoursCell);
    row.appendChild(reasonCell);
    row.appendChild(statusCell);
    row.appendChild(actionCell);

    return result;
  }

})();
