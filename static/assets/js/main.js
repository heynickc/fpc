/*
	Scalar by Pixelarity
	pixelarity.com @pixelarity
	License: pixelarity.com/license
*/

(function($) {

	skel.breakpoints({
		xlarge: '(max-width: 1680px)',
		large: '(max-width: 1280px)',
		medium: '(max-width: 980px)',
		small: '(max-width: 736px)',
		xsmall: '(max-width: 480px)'
	});

	$(function() {

		var	$window = $(window),
			$body = $('body'),
			$html = $('html');

		// Disable animations/transitions until the page has loaded.
			$body.addClass('is-loading');

			$window.on('load', function() {
				window.setTimeout(function() {
					$body.removeClass('is-loading');
				}, 0);
			});

		// Fix: Placeholder polyfill.
			$('form').placeholder();

		// Prioritize "important" elements on medium.
			skel.on('+medium -medium', function() {
				$.prioritize(
					'.important\\28 medium\\29',
					skel.breakpoint('medium').active
				);
			});

		// Dropdowns.
			$('#nav > ul').dropotron({
				alignment: 'center',
				hideDelay: 350
			});

		// Off-Canvas Navigation.

			// Navigation Button.
				$(
					'<div id="navButton">' +
						'<a href="#navPanel" class="toggle"></a>' +
					'</div>'
				)
					.appendTo($body);

			// Navigation Panel.
				$(
					'<div id="navPanel">' +
						'<nav>' +
							$('#nav').navList() +
						'</nav>' +
					'</div>'
				)
					.appendTo($body)
					.panel({
						delay: 500,
						hideOnClick: true,
						resetScroll: true,
						resetForms: true,
						side: 'top',
						target: $html,
						visibleClass: 'navPanel-visible'
					});

			// Fix: Remove transitions on WP<10 (poor/buggy performance).
				if (skel.vars.os == 'wp' && skel.vars.osVersion < 10)
					$('#header, #navButton, #navPanel, #page-wrapper')
						.css('transition', 'none');

		// IE8 fixes.
			if (skel.vars.IEVersion < 9) {

				// Odd/even on wrapper.split.
					$('.wrapper.split .secondary > section').each(function(i) {
						$(this).addClass((i + 1) % 2 == 0 ? 'even' : 'odd');
					});

			}

			var calendars = {};
			// assuming you've got the appropriate language files,
			// clndr will respect whatever moment's language is set to.
			// moment.locale('ru');

			// here's some magic to make sure the dates are happening this month.
			var thisMonth = moment().format('YYYY-MM');

			var eventArray = [
			{ startDate: thisMonth + '-10', endDate: thisMonth + '-14', title: 'Multi-Day Event' },
			{ startDate: thisMonth + '-21', endDate: thisMonth + '-23', title: 'Another Multi-Day Event' },
			{ date: thisMonth + '-27', title: 'Single Day Event' }
			];

			// the order of the click handlers is predictable.
			// direct click action callbacks come first: click, nextMonth, previousMonth, nextYear, previousYear, or today.
			// then onMonthChange (if the month changed).
			// finally onYearChange (if the year changed).

			calendars.clndr1 = $('.cal1').clndr({
			// events: eventArray,
			// constraints: {
			//   startDate: '2013-11-01',
			//   endDate: '2013-11-15'
			// },
			clickEvents: {
			  click: function(target) {
			    console.log(target);
			    // if you turn the `constraints` option on, try this out:
			    // if($(target.element).hasClass('inactive')) {
			    //   console.log('not a valid datepicker date.');
			    // } else {
			    //   console.log('VALID datepicker date.');
			    // }
			  },
			  nextMonth: function() {
			    console.log('next month.');
			  },
			  previousMonth: function() {
			    console.log('previous month.');
			  },
			  onMonthChange: function() {
			    console.log('month changed.');
			  },
			  nextYear: function() {
			    console.log('next year.');
			  },
			  previousYear: function() {
			    console.log('previous year.');
			  },
			  onYearChange: function() {
			    console.log('year changed.');
			  }
			},
			multiDayEvents: {
			  startDate: 'startDate',
			  endDate: 'endDate',
			  singleDay: 'date'
			},
			showAdjacentMonths: true,
			adjacentDaysChangeMonth: false
			});

	});

})(jQuery);