package stringHTML

//WallLoadStart ... unknown
func WallLoadStart(docID string, pageID string) string {
	return `
	<script>
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
                             $('.editEntry').click(function(e) {
                                var postId= $(this).attr("id");
                                var modified ="#"+postId;
                                $(modified).modal('show');
                             });
                            $('.comment-btn').click(function(e) {
                                        var postId= $(this).attr("id");
                                        var modified ="#"+postId;
                                   //     console.log(modified);
                                        var Pikachu = "` + docID + `";
                                        //e.preventDefault();
                                        $.ajax({
                                            type: 'GET',  
                                            url: 'http://ustart.today:5000/getComments/',
                                            contentType: "application/json; charset=utf-8",
                                            data: {userID:"` + pageID + `", PostID:postId,Pikachu:Pikachu},
                                            success: function(data) {
                                            	$("#commentnil").html(data);
                                             //   console.log(data);
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
                                                console.log("share clicked");
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
      
	</script>
	`
}
