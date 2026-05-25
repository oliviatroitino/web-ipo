// ── Menú móvil ──────────────────────────────────────────────────────────────
const navToggle = document.querySelector('.nav-toggle');
const mobileMenu = document.getElementById('mobile-menu');

if (navToggle && mobileMenu) {
  navToggle.addEventListener('click', () => {
    const isOpen = !mobileMenu.hidden;
    mobileMenu.hidden = isOpen;
    navToggle.setAttribute('aria-expanded', String(!isOpen));
  });

  // Cerrar el menú al hacer click en cualquier enlace dentro de él
  mobileMenu.querySelectorAll('a').forEach(link => {
    link.addEventListener('click', () => {
      mobileMenu.hidden = true;
      navToggle.setAttribute('aria-expanded', 'false');
    });
  });
}

// ── Toggle panel de filtros (móvil) ─────────────────────────────────────────
const filterToggle = document.querySelector('.js-filter-toggle');
const filterPanel  = document.getElementById('filter-panel');

if (filterToggle && filterPanel) {
  filterToggle.addEventListener('click', () => {
    const isOpen = filterPanel.classList.toggle('is-open');
    filterToggle.setAttribute('aria-expanded', String(isOpen));
  });
}

// ── Favoritos (localStorage) ─────────────────────────────────────────────────
// Se implementa en el paso 9. Por ahora exportamos las funciones base.

function getFavoritos() {
  try {
    return JSON.parse(localStorage.getItem('favoritos') || '[]');
  } catch {
    return [];
  }
}

function addFavorito(id) {
  const favs = getFavoritos();
  if (!favs.includes(id)) {
    favs.push(id);
    localStorage.setItem('favoritos', JSON.stringify(favs));
  }
}

function removeFavorito(id) {
  const favs = getFavoritos().filter(f => f !== id);
  localStorage.setItem('favoritos', JSON.stringify(favs));
}

function isFavorito(id) {
  return getFavoritos().includes(id);
}

// ── Enlace de favoritos en el nav ────────────────────────────────────────────
// Actualiza el href del enlace con los IDs actuales para que el servidor
// reciba la lista y pueda renderizar la página correctamente.
document.querySelectorAll('a[href="/favoritos"]').forEach(link => {
  const ids = getFavoritos();
  if (ids.length > 0) {
    link.href = '/favoritos?ids=' + ids.join(',');
  }
});

// ── Botón guardar en la página de detalle ────────────────────────────────────
const btnFavorito = document.querySelector('.js-favorito');

if (btnFavorito) {
  const id = btnFavorito.dataset.id;

  // Estado inicial
  if (isFavorito(id)) {
    btnFavorito.textContent = 'Guardado';
    btnFavorito.classList.add('btn--saved');
  }

  btnFavorito.addEventListener('click', () => {
    if (isFavorito(id)) {
      removeFavorito(id);
      btnFavorito.textContent = 'Guardar';
      btnFavorito.classList.remove('btn--saved');
    } else {
      addFavorito(id);
      btnFavorito.textContent = 'Guardado';
      btnFavorito.classList.add('btn--saved');
    }
  });
}
