/* eslint-disable @typescript-eslint/no-explicit-any */
/* eslint-disable react-hooks/set-state-in-effect */
import { useState, useEffect, useRef } from 'react';
import {
  Menu, X, ChevronDown, User, LogOut,
  LayoutDashboard, FileText, BarChart2, Settings,
  Users, Briefcase, ClipboardList, TrendingUp,
  Award, Shield,
} from 'lucide-react';
import logo from "/image/logo/logo-monalisa.png";
import { useNavigate, useLocation } from 'react-router-dom';
import { useAuth } from '../context/AuthContext';
import useAuthHooks from '../hooks/AuthHooks';
import { BASE_URL_FILE } from '../server/API';

type NavItem = { label: string; path: string; icon: React.ReactNode };

function NavButton({ label, icon, isActive, onClick }: {
  label: string; path: string; icon: React.ReactNode;
  isActive: boolean; onClick: () => void;
}) {
  return (
    <button
      onClick={onClick}
      className={`group flex items-center gap-2 px-4 py-2 rounded-xl text-sm transition-all duration-200
        ${isActive ? 'text-primary bg-primary/10 font-semibold' : 'text-gray-600 hover:text-primary hover:bg-primary/8'}`}
    >
      <span className={`transition-colors duration-200 ${isActive ? 'text-primary' : 'text-gray-400 group-hover:text-primary'}`}>
        {icon}
      </span>
      <span className="font-poppins-medium">{label}</span>
    </button>
  );
}

function DropdownMenu({ items, label, isOpen, onToggle, dropRef, isParentActive, icon, isActive, onNavigate }: {
  items: NavItem[]; label: string; isOpen: boolean; onToggle: () => void;
  dropRef: React.RefObject<HTMLDivElement | null>; paths: string[];
  icon: React.ReactNode; isParentActive: boolean;
  isActive: (path: string) => boolean; onNavigate: (path: string) => void;
}) {
  return (
    <div className="relative" ref={dropRef}>
      <button
        onClick={onToggle}
        className={`group flex items-center gap-2 px-4 py-2 rounded-xl text-sm transition-all duration-200
          ${isParentActive ? 'text-primary bg-primary/10' : 'text-gray-600 hover:text-primary hover:bg-primary/8'}`}
      >
        <span className={`transition-colors duration-200 ${isParentActive ? 'text-primary' : 'text-gray-400 group-hover:text-primary'}`}>
          {icon}
        </span>
        <span className="font-poppins-medium">{label}</span>
        <ChevronDown className={`h-3.5 w-3.5 transition-all duration-300 ${isOpen ? 'rotate-180 text-primary' : 'text-gray-400'}`} />
      </button>
      <div
        className={`absolute top-full left-0 mt-2 bg-white rounded-2xl shadow-2xl border border-gray-100 py-2 z-50 transition-all duration-200 origin-top-left
          ${isOpen ? 'opacity-100 scale-100 translate-y-0 pointer-events-auto' : 'opacity-0 scale-95 -translate-y-2 pointer-events-none'}`}
        style={{ minWidth: '260px' }}
      >
        <p className="text-[10px] font-semibold uppercase tracking-widest text-gray-400 px-5 py-2">{label}</p>
        {items.map((item, index) => (
          <button
            key={index}
            onClick={() => { onNavigate(item.path); onToggle(); }}
            className={`w-full flex items-center gap-3 px-5 py-2.5 text-sm transition-all duration-150
              ${isActive(item.path) ? 'text-primary bg-primary/10 font-semibold' : 'text-gray-600 hover:text-primary hover:bg-primary/8'}`}
          >
            <span className={`shrink-0 ${isActive(item.path) ? 'text-primary' : 'text-gray-400'}`}>{item.icon}</span>
            <span className="font-poppins-regular">{item.label}</span>
            {isActive(item.path) && <span className="ml-auto w-1.5 h-1.5 rounded-full bg-primary" />}
          </button>
        ))}
      </div>
    </div>
  );
}

function MobileDropdown({ label, items, isOpen, onToggle, icon, isActive, onNavigate }: {
  label: string; items: NavItem[]; isOpen: boolean; onToggle: () => void;
  icon: React.ReactNode; isActive: (path: string) => boolean; onNavigate: (path: string) => void;
}) {
  return (
    <div>
      <button
        onClick={onToggle}
        className={`w-full flex items-center justify-between px-4 py-3.5 rounded-2xl text-sm transition-all duration-150
          ${isOpen ? 'text-primary bg-primary/10' : 'text-gray-700 hover:text-primary hover:bg-primary/8'}`}
      >
        <span className="flex items-center gap-3">
          <span className={`p-1.5 rounded-lg transition-colors ${isOpen ? 'bg-primary/20 text-primary' : 'bg-gray-100 text-gray-500'}`}>
            {icon}
          </span>
          <span className="font-poppins-medium">{label}</span>
        </span>
        <ChevronDown className={`h-4 w-4 transition-all duration-300 ${isOpen ? 'rotate-180 text-primary' : 'text-gray-400'}`} />
      </button>
      <div className={`overflow-hidden transition-all duration-300 ease-in-out ${isOpen ? 'max-h-96 opacity-100' : 'max-h-0 opacity-0'}`}>
        <div className="ml-3 mt-1 mb-1 pl-3 border-l-2 border-primary/20 space-y-0.5">
          {items.map((item, index) => (
            <button
              key={index}
              onClick={() => onNavigate(item.path)}
              className={`w-full flex items-center gap-3 px-3 py-2.5 rounded-xl text-sm transition-all duration-150
                ${isActive(item.path) ? 'text-primary bg-primary/10 font-semibold' : 'text-gray-600 hover:text-primary hover:bg-primary/8'}`}
            >
              <span className={`${isActive(item.path) ? 'text-primary' : 'text-gray-400'}`}>{item.icon}</span>
              <span className="font-poppins-regular">{item.label}</span>
              {isActive(item.path) && <span className="ml-auto w-1.5 h-1.5 rounded-full bg-primary" />}
            </button>
          ))}
        </div>
      </div>
    </div>
  );
}

function UserAvatar({ user, size = 'md' }: { user: any; size?: 'sm' | 'md' }) {
  const dim = size === 'sm' ? 'w-8 h-8 text-xs' : 'w-10 h-10 text-sm';
  return (
    <div className={`${dim} rounded-full flex items-center justify-center shrink-0 ring-2 ring-white shadow-sm ${user?.file_photo ? '' : 'bg-linear-to-br from-primary/20 to-primary/50'}`}>
      {user?.file_photo ? (
        <img className={`rounded-full ${dim} object-cover`} src={`${BASE_URL_FILE}/${user.file_photo}`} alt="" />
      ) : (
        <span className="font-poppins-medium text-primary">
          {user?.fullname?.split(' ').map((n: string) => n[0]).join('').slice(0, 2)}
        </span>
      )}
    </div>
  );
}

function MobileUserSection({ user, type, onNavigate, onLogout }: {
  user: any; type: string; onNavigate: (path: string) => void; onLogout: () => void;
}) {
  return (
    <div className="border-t border-gray-100 pt-3 mt-2 space-y-1">
      <div className="flex items-center gap-3 px-4 py-3 bg-gray-50 rounded-2xl mb-2">
        <UserAvatar user={user} size="md" />
        <div className="flex-1 min-w-0">
          <p className="font-poppins-medium text-sm text-gray-900 truncate">{user?.fullname}</p>
          <span className="inline-block text-xs text-primary bg-primary/10 px-2 py-0.5 rounded-full font-poppins-medium mt-0.5">
            {type.toLocaleUpperCase()}
          </span>
        </div>
      </div>
      <button
        onClick={() => onNavigate('/ubah-profile')}
        className="w-full flex items-center gap-3 px-4 py-3 rounded-2xl text-sm text-gray-600 hover:text-primary hover:bg-primary/8 active:bg-primary/15 transition-all duration-150"
      >
        <span className="p-1.5 bg-gray-100 rounded-lg text-gray-500"><User className="h-4 w-4" /></span>
        <span className="font-poppins-regular">Edit Profile</span>
      </button>
      <button
        onClick={onLogout}
        className="w-full flex items-center gap-3 px-4 py-3 rounded-2xl text-sm text-red-500 hover:bg-red-50 active:bg-red-100 transition-all duration-150"
      >
        <span className="p-1.5 bg-red-50 rounded-lg text-red-400"><LogOut className="h-4 w-4" /></span>
        <span className="font-poppins-regular">Logout</span>
      </button>
    </div>
  );
}

export default function Navbar() {
  const [isMenuOpen, setIsMenuOpen] = useState<boolean>(false);
  const [isLaporanOpen, setIsLaporanOpen] = useState<boolean>(false);
  const [isHasilOpen, setIsHasilOpen] = useState<boolean>(false);
  const [isProfileOpen, setIsProfileOpen] = useState<boolean>(false);
  const [scrolled, setScrolled] = useState<boolean>(false);
  const profileRef = useRef<HTMLDivElement>(null);
  const laporanRef = useRef<HTMLDivElement>(null);
  const hasilRef = useRef<HTMLDivElement>(null);

  const navigate = useNavigate();
  const location = useLocation();
  const { user, loading } = useAuth();
  const { handleLogout } = useAuthHooks();
  const type = user ? user.role.name : 'guest';

  useEffect(() => {
    const handleScroll = () => setScrolled(window.scrollY > 10);
    window.addEventListener('scroll', handleScroll);
    return () => window.removeEventListener('scroll', handleScroll);
  }, []);

  useEffect(() => {
    const handleClickOutside = (e: MouseEvent) => {
      if (profileRef.current && !profileRef.current.contains(e.target as Node)) setIsProfileOpen(false);
      if (laporanRef.current && !laporanRef.current.contains(e.target as Node)) setIsLaporanOpen(false);
      if (hasilRef.current && !hasilRef.current.contains(e.target as Node)) setIsHasilOpen(false);
    };
    document.addEventListener('mousedown', handleClickOutside);
    return () => document.removeEventListener('mousedown', handleClickOutside);
  }, []);

  useEffect(() => {
    setIsMenuOpen(false);
    setIsLaporanOpen(false);
    setIsHasilOpen(false);
    setIsProfileOpen(false);
  }, [location.pathname]);

  const isActive = (path: string) => location.pathname === path;
  const isParentActive = (paths: string[]) => paths.some(p => location.pathname.startsWith(p));

  const ppkLaporanItems: NavItem[] = [
    { label: 'Rencana Anggaran', path: '/ppk/rencana-anggaran', icon: <FileText className="h-4 w-4" /> },
    { label: 'Jadwal Pelaksanaan', path: '/ppk/jadwal-pelaksanaan', icon: <ClipboardList className="h-4 w-4" /> },
    { label: 'Realisasi Pekerjaan', path: '/ppk/realisasi-pekerjaan', icon: <BarChart2 className="h-4 w-4" /> },
    { label: 'Project Progress (Kurva S)', path: '/ppk/project-kurva-s', icon: <TrendingUp className="h-4 w-4" /> },
  ];

  const pokjaLaporanItems: NavItem[] = [
    { label: 'Data Entry Penjabat Pengadaan', path: '/pokja/data-entry-penjabat-pengadaan', icon: <FileText className="h-4 w-4" /> },
    { label: 'Data Entry Kelompok Kerja', path: '/pokja/data-entry-kelompok-kerja', icon: <Briefcase className="h-4 w-4" /> },
  ];

  const pokjaHasilItems: NavItem[] = [
    { label: 'Penjabat Pengadaan', path: '/pokja/penjabat-pengadaan', icon: <Award className="h-4 w-4" /> },
    { label: 'Kelompok Kerja', path: '/pokja/kelompok-kerja', icon: <Users className="h-4 w-4" /> },
  ];

  const kepalaLaporanItems: NavItem[] = [
    { label: 'Rencana Anggaran', path: '/kepala/rencana-anggaran', icon: <FileText className="h-4 w-4" /> },
    { label: 'Jadwal Pelaksanaan', path: '/kepala/jadwal-pelaksanaan', icon: <ClipboardList className="h-4 w-4" /> },
    { label: 'Realisasi Pekerjaan', path: '/kepala/realisasi-pekerjaan', icon: <BarChart2 className="h-4 w-4" /> },
    { label: 'Project Progress (Kurva S)', path: '/kepala/project-kurva-s', icon: <TrendingUp className="h-4 w-4" /> },
  ];

  const kepalaHasilItems: NavItem[] = [
    { label: 'Penjabat Pengadaan', path: '/kepala/penjabat-pengadaan', icon: <Award className="h-4 w-4" /> },
    { label: 'Kelompok Kerja', path: '/kepala/kelompok-kerja', icon: <Users className="h-4 w-4" /> },
  ];

  const adminItems: NavItem[] = [
    { label: 'Manajemen Pengguna', path: '/admin/manajemen-pengguna', icon: <Users className="h-4 w-4" /> },
    { label: 'Kelompok Kerja', path: '/admin/kelompok-kerja', icon: <Briefcase className="h-4 w-4" /> },
  ];

  if (loading) return null;

  return (
    <nav className={`w-full fixed top-0 left-0 right-0 z-50 transition-all duration-300
      ${scrolled ? 'bg-white/95 backdrop-blur-md shadow-lg shadow-gray-200/60 border-b border-gray-100' : 'bg-white shadow-sm'}
      ${type === 'guest' ? 'py-3 md:py-4' : 'py-2 md:py-2'}`}
    >
      <div className="max-w-7xl mx-auto px-4 md:px-8 flex items-center justify-between">
        <div className="flex items-center gap-6">
          <button onClick={() => navigate("/")} className="group">
            <img src={logo} className="w-auto h-7 transition-transform duration-200 group-hover:scale-105" />
          </button>

          {type === 'ppk' && (
            <div className="hidden md:flex items-center gap-1">
              <NavButton label="Dashboard" path="/" icon={<LayoutDashboard className="h-4 w-4" />} isActive={isActive("/")} onClick={() => navigate("/")} />
              <DropdownMenu label="Laporan Saya" items={ppkLaporanItems} isOpen={isLaporanOpen}
                onToggle={() => setIsLaporanOpen(!isLaporanOpen)} dropRef={laporanRef}
                paths={['/ppk/']} icon={<FileText className="h-4 w-4" />}
                isParentActive={isParentActive(['/ppk/'])} isActive={isActive} onNavigate={navigate} />
            </div>
          )}

          {type === 'pokja/pp' && (
            <div className="hidden md:flex items-center gap-1">
              <NavButton label="Dashboard" path="/" icon={<LayoutDashboard className="h-4 w-4" />} isActive={isActive("/")} onClick={() => navigate("/")} />
              <DropdownMenu label="Laporan Saya" items={pokjaLaporanItems} isOpen={isLaporanOpen}
                onToggle={() => { setIsLaporanOpen(!isLaporanOpen); setIsHasilOpen(false); }} dropRef={laporanRef}
                paths={['/pokja/data-entry']} icon={<FileText className="h-4 w-4" />}
                isParentActive={isParentActive(['/pokja/data-entry'])} isActive={isActive} onNavigate={navigate} />
              <DropdownMenu label="Laporan Hasil Pemilihan" items={pokjaHasilItems} isOpen={isHasilOpen}
                onToggle={() => { setIsHasilOpen(!isHasilOpen); setIsLaporanOpen(false); }} dropRef={hasilRef}
                paths={['/pokja/penjabat', '/pokja/kelompok']} icon={<Award className="h-4 w-4" />}
                isParentActive={isParentActive(['/pokja/penjabat', '/pokja/kelompok'])} isActive={isActive} onNavigate={navigate} />
            </div>
          )}

          {(type === 'kepala bagian' || type === 'kepala biro') && (
            <div className="hidden md:flex items-center gap-1">
              <NavButton label="Dashboard" path="/" icon={<LayoutDashboard className="h-4 w-4" />} isActive={isActive("/")} onClick={() => navigate("/")} />
              <DropdownMenu label="Laporan Saya" items={kepalaLaporanItems} isOpen={isLaporanOpen}
                onToggle={() => { setIsLaporanOpen(!isLaporanOpen); setIsHasilOpen(false); }} dropRef={laporanRef}
                paths={['/kepala/rencana', '/kepala/jadwal', '/kepala/realisasi', '/kepala/project']} icon={<FileText className="h-4 w-4" />}
                isParentActive={isParentActive(['/kepala/rencana', '/kepala/jadwal', '/kepala/realisasi', '/kepala/project'])} isActive={isActive} onNavigate={navigate} />
              <DropdownMenu label="Laporan Hasil Pemilihan" items={kepalaHasilItems} isOpen={isHasilOpen}
                onToggle={() => { setIsHasilOpen(!isHasilOpen); setIsLaporanOpen(false); }} dropRef={hasilRef}
                paths={['/kepala/penjabat', '/kepala/kelompok']} icon={<Award className="h-4 w-4" />}
                isParentActive={isParentActive(['/kepala/penjabat', '/kepala/kelompok'])} isActive={isActive} onNavigate={navigate} />
            </div>
          )}

          {type === 'admin' && (
            <div className="hidden md:flex items-center gap-1">
              <NavButton label="Dashboard" path="/" icon={<LayoutDashboard className="h-4 w-4" />} isActive={isActive("/")} onClick={() => navigate("/")} />
              <NavButton label="Pengaturan API" path="/admin/konfigurasi" icon={<Settings className="h-4 w-4" />} isActive={isActive("/admin/konfigurasi")} onClick={() => navigate("/admin/konfigurasi")} />
              <DropdownMenu label="Manajemen" items={adminItems} isOpen={isLaporanOpen}
                onToggle={() => setIsLaporanOpen(!isLaporanOpen)} dropRef={laporanRef}
                paths={['/admin/']} icon={<Shield className="h-4 w-4" />}
                isParentActive={isParentActive(['/admin/'])} isActive={isActive} onNavigate={navigate} />
            </div>
          )}
        </div>

        {type === 'guest' && (
          <>
            <div className="hidden md:flex">
              <button onClick={() => navigate("/masuk")}
                className="flex items-center gap-2 px-5 py-2.5 rounded-xl bg-primary text-white text-sm font-semibold shadow-md shadow-primary/30 hover:opacity-90 active:scale-95 transition-all duration-200"
              >
                <User className="h-4 w-4" />
                <span className="font-poppins-medium">Masuk</span>
              </button>
            </div>
            <button className="md:hidden p-2.5 rounded-xl hover:bg-primary/10 active:bg-primary/20 transition-all" onClick={() => setIsMenuOpen(prev => !prev)}>
              {isMenuOpen ? <X className="h-5 w-5 text-primary" /> : <Menu className="h-5 w-5 text-gray-600" />}
            </button>
          </>
        )}

        {type !== 'guest' && (
          <>
            <div className="hidden md:flex items-center">
              <div className="relative" ref={profileRef}>
                <button
                  onClick={() => setIsProfileOpen(!isProfileOpen)}
                  className={`flex items-center gap-2.5 pl-2 pr-3 py-1.5 rounded-xl transition-all duration-200 ${isProfileOpen ? 'bg-primary/10' : 'hover:bg-gray-100'}`}
                >
                  <UserAvatar user={user} size="sm" />
                  <div className="text-left hidden lg:block">
                    <p className="font-poppins-medium text-xs text-gray-900 leading-tight">{user?.fullname}</p>
                    <p className="text-[10px] text-gray-400 capitalize leading-tight">{type}</p>
                  </div>
                  <ChevronDown className={`h-3.5 w-3.5 text-gray-400 transition-all duration-300 ${isProfileOpen ? 'rotate-180 text-primary' : ''}`} />
                </button>
                <div className={`absolute top-full right-0 mt-2 w-56 bg-white rounded-2xl shadow-2xl border border-gray-100 py-2 z-50 transition-all duration-200 origin-top-right
                  ${isProfileOpen ? 'opacity-100 scale-100 translate-y-0 pointer-events-auto' : 'opacity-0 scale-95 -translate-y-2 pointer-events-none'}`}
                >
                  <div className="px-4 py-3 border-b border-gray-50">
                    <p className="font-poppins-medium text-sm text-gray-900">{user?.fullname}</p>
                    <p className="text-xs text-gray-400 capitalize mt-0.5">{type.toLocaleUpperCase()}</p>
                  </div>
                  <div className="pt-1.5">
                    <button onClick={() => { navigate("/ubah-profile"); setIsProfileOpen(false); }}
                      className="w-full flex items-center gap-3 px-4 py-2.5 text-sm text-gray-600 hover:text-primary hover:bg-primary/8 transition-all duration-150"
                    >
                      <User className="h-4 w-4 text-gray-400" />
                      <span className="font-poppins-regular">Ubah Profile</span>
                    </button>
                    <button onClick={() => { handleLogout(); setIsProfileOpen(false); }}
                      className="w-full flex items-center gap-3 px-4 py-2.5 text-sm text-red-500 hover:bg-red-50 transition-all duration-150"
                    >
                      <LogOut className="h-4 w-4" />
                      <span className="font-poppins-regular">Logout</span>
                    </button>
                  </div>
                </div>
              </div>
            </div>
            <button className="md:hidden p-2.5 rounded-xl hover:bg-primary/10 active:bg-primary/20 transition-all" onClick={() => setIsMenuOpen(prev => !prev)}>
              {isMenuOpen ? <X className="h-5 w-5 text-primary" /> : <Menu className="h-5 w-5 text-gray-600" />}
            </button>
          </>
        )}
      </div>

      {isMenuOpen && (
        <div className="md:hidden border-t border-gray-100 mt-2">
          <div className="px-4 pt-3 pb-6 space-y-1.5">
            {type === 'guest' && (
              <button onClick={() => navigate("/masuk")}
                className="w-full flex items-center justify-center gap-2 px-4 py-3.5 rounded-2xl bg-primary text-white text-sm font-semibold shadow-md shadow-primary/20 active:opacity-90 transition-all"
              >
                <User className="h-4 w-4" />
                <span className="font-poppins-medium">Masuk ke Akun</span>
              </button>
            )}
            {type === 'ppk' && (<>
              <NavButton label="Dashboard" path="/" icon={<LayoutDashboard className="h-4 w-4" />} isActive={isActive("/")} onClick={() => navigate("/")} />
              <MobileDropdown label="Laporan Saya" items={ppkLaporanItems} isOpen={isLaporanOpen} onToggle={() => setIsLaporanOpen(!isLaporanOpen)} icon={<FileText className="h-4 w-4" />} isActive={isActive} onNavigate={navigate} />
              <MobileUserSection user={user} type={type} onNavigate={navigate} onLogout={handleLogout} />
            </>)}
            {type === 'pokja/pp' && (<>
              <NavButton label="Dashboard" path="/" icon={<LayoutDashboard className="h-4 w-4" />} isActive={isActive("/")} onClick={() => navigate("/")} />
              <MobileDropdown label="Laporan Saya" items={pokjaLaporanItems} isOpen={isLaporanOpen} onToggle={() => { setIsLaporanOpen(!isLaporanOpen); setIsHasilOpen(false); }} icon={<FileText className="h-4 w-4" />} isActive={isActive} onNavigate={navigate} />
              <MobileDropdown label="Laporan Hasil Pemilihan" items={pokjaHasilItems} isOpen={isHasilOpen} onToggle={() => { setIsHasilOpen(!isHasilOpen); setIsLaporanOpen(false); }} icon={<Award className="h-4 w-4" />} isActive={isActive} onNavigate={navigate} />
              <MobileUserSection user={user} type={type} onNavigate={navigate} onLogout={handleLogout} />
            </>)}
            {(type === 'kepala bagian' || type === 'kepala biro') && (<>
              <NavButton label="Dashboard" path="/" icon={<LayoutDashboard className="h-4 w-4" />} isActive={isActive("/")} onClick={() => navigate("/")} />
              <MobileDropdown label="Laporan Saya" items={kepalaLaporanItems} isOpen={isLaporanOpen} onToggle={() => { setIsLaporanOpen(!isLaporanOpen); setIsHasilOpen(false); }} icon={<FileText className="h-4 w-4" />} isActive={isActive} onNavigate={navigate} />
              <MobileDropdown label="Laporan Hasil Pemilihan" items={kepalaHasilItems} isOpen={isHasilOpen} onToggle={() => { setIsHasilOpen(!isHasilOpen); setIsLaporanOpen(false); }} icon={<Award className="h-4 w-4" />} isActive={isActive} onNavigate={navigate} />
              <MobileUserSection user={user} type={type} onNavigate={navigate} onLogout={handleLogout} />
            </>)}
            {type === 'admin' && (<>
              <NavButton label="Dashboard" path="/" icon={<LayoutDashboard className="h-4 w-4" />} isActive={isActive("/")} onClick={() => navigate("/")} />
              <NavButton label="Pengaturan API" path="/admin/konfigurasi" icon={<Settings className="h-4 w-4" />} isActive={isActive("/admin/konfigurasi")} onClick={() => navigate("/admin/konfigurasi")} />
              <MobileDropdown label="Manajemen" items={adminItems} isOpen={isLaporanOpen} onToggle={() => setIsLaporanOpen(!isLaporanOpen)} icon={<Shield className="h-4 w-4" />} isActive={isActive} onNavigate={navigate} />
              <MobileUserSection user={user} type={type} onNavigate={navigate} onLogout={handleLogout} />
            </>)}
          </div>
        </div>
      )}
    </nav>
  );
}