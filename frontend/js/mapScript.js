(() => {
    const mapContainer = document.getElementById('map');
    if (!mapContainer) {
        return;
    }

    if (typeof mapboxgl === "undefined") {
        console.error("Mapbox GL JS has not been loaded.");
        return;
    }

    const token = window.MAPBOX_ACCESS_TOKEN || "";
    if (!token) {
        mapContainer.innerHTML = "Map unavailable. Set MAPBOX_ACCESS_TOKEN on the server.";
        mapContainer.classList.add("mapbox-token-missing");
        return;
    }

    mapboxgl.accessToken = token;

    const map = new mapboxgl.Map({
        container: 'map',
        style: 'mapbox://styles/mapbox/streets-v12',
        center: [0, 25],
        zoom: 1.3
    });

    map.addControl(new mapboxgl.NavigationControl());
    map.addControl(new mapboxgl.FullscreenControl());

    const bounds = new mapboxgl.LngLatBounds();
    let hasMarkers = false;

    const addMarker = (location, coordinates) => {
        new mapboxgl.Marker({ color: '#ff3636' })
            .setLngLat(coordinates)
            .setPopup(new mapboxgl.Popup({ offset: 25 }).setText(location))
            .addTo(map);
        bounds.extend(coordinates);
        hasMarkers = true;
    };

    const geocodeLocation = async (location) => {
        const url = `https://api.mapbox.com/geocoding/v5/mapbox.places/${encodeURIComponent(location)}.json?access_token=${token}&limit=1`;
        try {
            const response = await fetch(url);
            if (!response.ok) {
                throw new Error(`Failed to geocode ${location}: ${response.status}`);
            }
            const data = await response.json();
            if (data.features && data.features.length > 0) {
                addMarker(location, data.features[0].geometry.coordinates);
            }
        } catch (error) {
            console.error(error);
        }
    };

    const locationsList = Array.isArray(window.locations) ? window.locations : [];
    Promise.all(locationsList.map(geocodeLocation)).then(() => {
        if (hasMarkers) {
            map.fitBounds(bounds, { padding: 60, maxZoom: 5 });
        }
    });
})();
