$(document).ready(function () {
	$('.mob__nav').on('click', function () {
		$(this).slideDown();
		$('.nav').slideToggle('active');
	});
	$('.team_slider').slick({
		dots: true,
		speed: 300,
		slidesToShow: 1,
		adaptiveHeight: true,
		responsive: [
			{
				breakpoint: 768,
				settings: {
					arrows: false,
					slidrToShow: 2
				}
    }
  ]
	});
	$('.photo_slider').slick({
		dots: true,
		speed: 300,
		slidesToShow: 4,
		adaptiveHeight: true,
		arrows: true,
		responsive: [
			{
				breakpoint: 768,
				settings: {
					arrows: false,
					slidrToShow: 2
				}
    }
  ]
	});
});
$('.post_slider').slick({
	dots: true,
	speed: 300,
	slidesToShow: 1,
	adaptiveHeight: true,
	arrows: false
});