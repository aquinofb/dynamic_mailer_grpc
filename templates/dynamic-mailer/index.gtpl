<html>
    <head>
      <title></title>
      <link href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/css/bootstrap.min.css" rel="stylesheet" integrity="sha384-BVYiiSIFeK1dGmJRAkycuHAHRg32OmUcww7on3RYdg4Va+PmSTsz/K68vbdEjh4u" crossorigin="anonymous">

      <style type="text/css">
        body {
          background-image: linear-gradient(-180deg,#eee,#fbfbfb);
        }

        .panel {
          background: #fff;
          box-shadow: 0 1px 3px 0 rgba(0,0,0,.25);
          box-sizing: border-box;
        }

        input[type=submit] {
          background-color: #9a6dc8;
          border-color: #8659b4;
        }

      </style>
    </head>
    <body class="container">
      <div style="position: absolute; top: -3px; width: 71px; overflow: hidden; right: 40px;">
        <img src="https://engineroom.teamwork.com/wp-content/uploads/2014/08/b-2.png" style="width: 140px; position: relative; left: -66px;">
      </div>

      <br />
      <br />
      <br />
      <div class="panel col-md-6 col-md-offset-3">
        <h3>Good Stuff!!!</h3>

        <hr>

        <form action="/dynamic-mailer" method="post" class="form">
          <div class="form-group">
            <label class="label-control">Consumer Network UUID:</label>
            <input type="text" name="message[reference_uuid]" class="form-control">
          </div>
          <div class="form-group">
            <label>Event:</label>
            <input type="text" name="message[event]" class="form-control">
          </div>
          <div class="form-group">
            <label>Email:</label>
            <input type="text" name="message[email]" class="form-control">
          </div>

          <div class="form-group">
            <img src="https://36b2b52ed7ec81e2d24c-af98f47f6df85e5d1c0cc9be25d59a29.ssl.cf3.rackcdn.com/logos/cropped.jpg?v=31" alt="babylon">
            <input type="submit" class="btn btn-primary pull-right" value="Send Dynamic Mailer">
          </div>

          <div class="clearfix"></div>
        </form>
      </div>
    </body>
</html>
