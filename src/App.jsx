
import React, { useEffect, useRef } from 'react';
import * as THREE from 'three';
import { GLTFLoader } from 'three/examples/jsm/loaders/GLTFLoader.js';
import './App.css';

function App() {
  const mountRef = useRef(null);

  useEffect(() => {
    const scene = new THREE.Scene();
    // Placeholder camera, will be replaced by GLB cameras if available
    let camera = new THREE.PerspectiveCamera(
      75,
      window.innerWidth / window.innerHeight,
      0.1,
      1000
    );
    const renderer = new THREE.WebGLRenderer({ antialias: true });
  renderer.setSize(window.innerWidth, window.innerHeight);
  renderer.domElement.style.display = 'block';
  renderer.domElement.style.width = '100vw';
  renderer.domElement.style.height = '100vh';
  mountRef.current.style.width = '100vw';
  mountRef.current.style.height = '100vh';
  mountRef.current.style.overflow = 'hidden';
  mountRef.current.style.margin = '0';
  mountRef.current.appendChild(renderer.domElement);


  // Ambient light (boosted)
  const ambientLight = new THREE.AmbientLight(0xffffff, 0.8);
  scene.add(ambientLight);

  // Hemisphere light for global fill
  const hemiLight = new THREE.HemisphereLight(0xffffff, 0x444444, 0.6);
  hemiLight.position.set(0, 20, 0);
  scene.add(hemiLight);

  // Directional light for stronger lighting and shadows
  const dirLight = new THREE.DirectionalLight(0xffffff, 1);
  dirLight.position.set(5, 10, 7);
  dirLight.castShadow = true;
  scene.add(dirLight);

  // Debug helpers (optional, comment out if not needed)
  // scene.add(new THREE.DirectionalLightHelper(dirLight, 5));
  // scene.add(new THREE.HemisphereLightHelper(hemiLight, 5));

    // load model
    let loadedScene = null;

    const loader = new GLTFLoader();
    let gltfCameras = [];
    let activeCameraIndex = 0;
    loader.load(
      '/models/office.glb',
      (gltf) => {
        loadedScene = gltf.scene;
        scene.add(loadedScene);
        // Find cameras in the GLB
        if (gltf.cameras && gltf.cameras.length > 0) {
          gltfCameras = gltf.cameras;
          camera = gltfCameras[0];
          camera.aspect = window.innerWidth / window.innerHeight;
          camera.updateProjectionMatrix();
          console.log('Using GLB camera:', camera.name || 0);
        } else {
          // fallback to default camera
          camera.position.set(0, 2, 5);
        }
        // Save to ref for switching
        window.__gltfCameras = gltfCameras;
        window.__setActiveCamera = (idx) => {
          if (gltfCameras[idx]) {
            camera = gltfCameras[idx];
            camera.aspect = window.innerWidth / window.innerHeight;
            camera.updateProjectionMatrix();
            activeCameraIndex = idx;
            console.log('Switched to camera', idx, camera.name);
          }
        };
      },
      undefined,
      (error) => {
        console.error('Error loading GLB:', error);
      }
    );


  // camera.position.set(0, 2, 5); // Only if not using GLB camera

    // Handle window resize
    const handleResize = () => {
      camera.aspect = window.innerWidth / window.innerHeight;
      camera.updateProjectionMatrix();
      renderer.setSize(window.innerWidth, window.innerHeight);
      renderer.domElement.style.width = '100vw';
      renderer.domElement.style.height = '100vh';
      if (mountRef.current) {
        mountRef.current.style.width = '100vw';
        mountRef.current.style.height = '100vh';
      }
    };
    window.addEventListener('resize', handleResize);


    function animate() {
      requestAnimationFrame(animate);
      renderer.render(scene, camera);
    }
    animate();

    // Switch camera on click
    const handleClick = () => {
      if (gltfCameras.length > 1) {
        activeCameraIndex = (activeCameraIndex + 1) % gltfCameras.length;
        camera = gltfCameras[activeCameraIndex];
        camera.aspect = window.innerWidth / window.innerHeight;
        camera.updateProjectionMatrix();
        console.log('Switched to camera', activeCameraIndex, camera.name);
      }
    };
    window.addEventListener('click', handleClick);

    // Clean up on unmount
    return () => {
      window.removeEventListener('resize', handleResize);
      window.removeEventListener('click', handleClick);
      renderer.dispose();
      if (loadedScene) {
        scene.remove(loadedScene);
        loadedScene.traverse((child) => {
          if (child.isMesh) {
            if (child.geometry) child.geometry.dispose();
            if (child.material) {
              if (Array.isArray(child.material)) {
                child.material.forEach((m) => m.dispose && m.dispose());
              } else {
                child.material.dispose && child.material.dispose();
              }
            }
            if (child.texture) child.texture.dispose && child.texture.dispose();
          }
        });
      }
      if (mountRef.current && renderer.domElement.parentNode === mountRef.current) {
        mountRef.current.removeChild(renderer.domElement);
      }
    };
  }, []);

  // Prevent scrolling on body
  useEffect(() => {
    document.body.style.overflow = 'hidden';
    document.body.style.margin = '0';
    document.body.style.padding = '0';
    return () => {
      document.body.style.overflow = '';
      document.body.style.margin = '';
      document.body.style.padding = '';
    };
  }, []);

  return (
    <div ref={mountRef} style={{ width: '100vw', height: '100vh', overflow: 'hidden', margin: 0, padding: 0 }} />
  );
}

export default App;
