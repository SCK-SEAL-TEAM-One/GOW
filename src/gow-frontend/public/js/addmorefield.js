$(function() {
  $("#addMore").click(function(e) {
    e.preventDefault();
    var $newfield = $(
      "<tbody><tr>" +
        "<td></td>" +
        "<td><div class='col-sm-12 col-md-12'><textarea class='form-control' id='description' name='description[]' rows='5' placeholder='รายการ'></textarea></div></td>" +
        "<td><div class='col-sm-12'><input type='text' id='amount' class='form-control' name='amount[]' placeholder='จำนวน'></div></td>" +
        "<td><div class='col-sm-12'><input type='text' id='pricePerUnit' class='form-control' name='pricePerUnit[]' placeholder='ราคาต่อหน่วย'></div></td>" +
        "<td><div class='col-sm-12'><input type='text' id='price' class='form-control' name='price[]' placeholder='จำนวนเงิน'></div></td>" +
        "</tr></tbody>"
    );
    $("#fieldList").append($newfield);
  });
});
