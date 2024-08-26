 // Set the Cesium access token
 Cesium.Ion.defaultAccessToken = "{{.Token}}"; // Initialize the Cesium viewer
 const viewer = new Cesium.Viewer('cesiumContainer');
 // Check if there are coordinates to display
 const lat = parseFloat("{{.Lat}}");
    const lon = parseFloat("{{.Lon}}");
 console.log(lat,lon)
 if (lat && lon) {
     viewer.camera.flyTo({
         destination: Cesium.Cartesian3.fromDegrees(lon, lat, 150000)
     });
 }