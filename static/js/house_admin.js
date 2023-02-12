function debounce(fn,delay){
    let delays=delay||1000;
    let timer;
    return function(){
        let th=this;
        let args=arguments;
        if (timer) {
            clearTimeout(timer);
        }
        timer=setTimeout(function () {
                timer=null;
                fn.apply(th,args);
        }, delays);
    };
}


$(function () {
    if (document.location.href.indexOf("/add/") !== -1){
        $('.field-province').css('display', "none");
        $('.field-city').css('display', "none");
        $('.field-longitude').css('display', "none");
        $('.field-region').css('display', "none");
        $('.field-latitude').css('display', "none");
    }

    $("#id_address").on('input propertychange',debounce(function(){
        if ($('#select-address').length <= 0) {
            $("#id_address").after("<select id='select-address'>请选择地址</select>");
            $("#select-address").change(function () {
                var address_obj = $(this).find(":selected").data("value");
                $('input[name="address"]').val(address_obj.title);
                // set_none_or_block("block")
                $('input[name="city"]').val(address_obj.city);
                $('input[name="region"]').val(address_obj.district);
                $('input[name="province"]').val(address_obj.province);
                $('input[name="longitude"]').val(address_obj.location.lng);
                $('input[name="latitude"]').val(address_obj.location.lat);
                // set_none_or_block("none")
            });
        }
        var address = $('#id_address').val();
        var $el = $("#select-address")
        $el.empty();
        if (address) {
            $.ajax({
                async: true,
                type: 'GET',
                dataType: 'jsonp',
                url: 'https://apis.map.qq.com/ws/place/v1/suggestion',
                data: {
                    keyword: address,
                    key: 'PZ7BZ-HMBLQ-WML5R-GW37R-6IW3Q-IQFNZ',
                    output: "jsonp",
                    page_size: 10,
                    page_index: 1,
                },
                success: function (resp) {
                    if(resp.status === 0 ){
                       $.each(resp.data, function(index , cnf) {
                           var json_value = JSON.stringify(cnf);
                           var address = `[${cnf.city}][${cnf.district}] ** [${cnf.title}]  (${cnf.address})`
                           $el.append($("<option></option>").attr("data-value", json_value).text(
                               address
                           ))
                       });
                       $el.change();
                    }
                }, error: function (data) {
                    alert(data.message)
                }
            });
        }
    }))
})