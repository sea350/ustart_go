package stringHTML

//WallLoadEnd ... unknown
func WallLoadEnd(docID string, pageID string) string {
	return `
	<script>
	// write new deleteEntry here after it works 
    // $('.deleteEntry').click(function(e) {
    //     //e.preventDefault();
    //     var postId= $(this).attr("id");
    //     var modified ="#"+postId;
    //     console.log(modified);
    //     $.ajax({
    //         type: 'GET',  
    //         url: 'http://ustart.today:5000/deleteModal/',
    //         contentType: "application/json; charset=utf-8",
    //         data: {},
    //         success: function(data) {
    //             console.log("reached here");
    //             $("#deletenil").html(data);
    //             $(modified).modal('show');
         
    //         }//end success
    //     });
    // });

		
		$('#commentform').submit(function(e) {
			//e.preventDefault();
			var $form = $(this);
			var docID = "` + docID + `";
			var text = $form.find('input[name=commentz]').val();
			console.log(text);
			$.ajax({
				type: 'GET',  
				url: 'http://ustart.today:5000/AddComment/',
				contentType: "application/json; charset=utf-8",
				data: {docID:docID,text:text},
				success: function(data) {
				//	console.log(data);
					$('.comments-list').append(data);
				}
			});
		});

		$('#new-postSubmit').click(function(e) {
			//e.preventDefault();
			var docID = "` + docID + `";
			var text = $("#post-msg").val();
			console.log(text);
			$.ajax({
				type: 'GET',  
				url: 'http://ustart.today:5000/addPost/',
				contentType: "application/json; charset=utf-8",
				data: {docID:docID,text:text},
				success: function(data) {
				//	console.log(data);
					$(".wallAppend").prepend(data);
					console.log('hello m8');
												 $(".comment-btn").hover(function (e) {
						var cmtBtnImg = $(this).find('img');
						cmtBtnImg.attr('src', "/ustart_front/ico/comment.png");     
					 },function (e) {
						var cmtBtnImg = $(this).find('img');
						cmtBtnImg.attr('src', "/ustart_front/ico/no comment.png");     
					 });   
					$(".share-btn").hover(function (e) {
						var shrBtnImg = $(this).find('img');
						shrBtnImg.attr('src', "/ustart_front/ico/share.png");     
					 },function (e) {
						var shrBtnImg = $(this).find('img');
						shrBtnImg.attr('src', "/ustart_front/ico/not share.png");     
					 });
					  $(".like-btn").hover(function (e) {
						var likeBtnImg = $(this).find('img');
						if (likeBtnImg.attr('src') === "/ustart_front/ico/like.png") {
							likeBtnImg.attr('src', "/ustart_front/ico/liked.png");
						} else {
							likeBtnImg.attr('src', "/ustart_front/ico/like.png");
						}
						return false;
					});
					$(".comment-like").click(function (e) {
						if ($(this).html() == "Like") {
							$(this).html('Liked');
						} else {
							$(this).html('Like');
						}
						return false;
					});
					$('body').on('click', '.odom-submit', function (e) {
						$('#shareCommentForm').submit();
					});
					$('.comment-btn').click(function(e) {
						var postId= $(this).attr("id");
						var modified ="#"+postId;
						console.log(modified);
						var Pikachu = "` + docID + `";
						//e.preventDefault();
						$.ajax({
							type: 'GET',  
							url: 'http://ustart.today:5000/getComments/',
							contentType: "application/json; charset=utf-8",
							data: {userID:"` + pageID + `", PostID:postId,Pikachu:Pikachu},
							success: function(data) {
								$("#commentnil").html(data);
								console.log(data);
								$(modified).modal('show');
							}
						});
					});
			 
					$('.share-btn').click(function(e) {
						var postId= $(this).attr("id");
						var modified ="#"+postId;
						console.log(modified);
						var Pikachu = "` + docID + `";
						//e.preventDefault();
						$.ajax({
							type: 'GET',  
							url: 'http://ustart.today:5000/shareComments/',
							contentType: "application/json; charset=utf-8",
							data: {userID:"` + pageID + `", PostID:postId,Pikachu:Pikachu},
							success: function(data) {
								$("#sharenil").html(data);
								console.log("share clicked ");
								$(modified).modal('show');
							}
						});
					});  
					$('.like-btn').click(function(e) {
						var postId= $(this).attr("id");
						var modified ="#"+postId;
						console.log(modified);
						var selfDoc = "` + docID + `";
						//e.preventDefault();
						$.ajax({
							type: 'GET',  
							url: 'http://ustart.today:5000/Like',
							contentType: "application/json; charset=utf-8",
							data: {userID:"` + pageID + `", PostID:postId,selfDoc:selfDoc},
							success: function(data) {
									var likeBtnImg = $(this).find('img');
									if (likeBtnImg.attr('src') === "/ustart_front/ico/like.png") {
										likeBtnImg.attr('src', "/ustart_front/ico/liked.png");
									} else {
										likeBtnImg.attr('src', "/ustart_front/ico/like.png");
									}
								console.log("like clicked");
							}
						});
					}); 
			   	}
			});
		});      

	</script> `
}
