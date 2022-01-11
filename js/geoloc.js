

    let arr = [];
    let el = document.getElementsByClassName("pDl")
    

    // for (i = 0; i < el.length; i++) {
    //     var myGeocoder = ymaps.geocode(el[i].innerHTML);
    //     myGeocoder.then(function(res) {
    //         myMap.geoObjects.add(res.geoObjects);
    //     });
    // }
    ymaps.ready(init);

    function init() {
        var myMap = new ymaps.Map('YMapsID', {
            center: [55.753994, 37.622093],
            zoom: 15
        });
    
        // Поиск координат центра Нижнего Новгорода.
        for (i = 0; i < el.length; i++) {
            let text = el[i].innerHTML
            let res1 = text.replace(/_/g, " ");
            let res2 = res1.replace(/-/g, " ");
            ymaps.geocode(res2, {
                results: 1
            }).then(function (res) {
                    var firstGeoObject = res.geoObjects.get(0),
                        coords = firstGeoObject.geometry.getCoordinates(),
                        bounds = firstGeoObject.properties.get('boundedBy');
                    firstGeoObject.options.set('preset', 'islands#darkBlueDotIconWithCaption');
                    firstGeoObject.properties.set('iconCaption', firstGeoObject.getAddressLine());
                    myMap.geoObjects.add(firstGeoObject);
                    myMap.setBounds(bounds, {
                        checkZoomRange: true
                    });
                });
        }
        
    }