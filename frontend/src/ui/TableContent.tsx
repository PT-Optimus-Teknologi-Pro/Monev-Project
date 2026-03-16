/* eslint-disable @typescript-eslint/no-explicit-any */
import { Eye, Edit2, CheckCircle, ChevronLeft, ChevronRight, FileX, LayoutList } from 'lucide-react';
import { useState } from 'react';

interface TableColumn {
    key: string;
    label: string;
}

interface TableContentProps {
    columns: TableColumn[];
    data: any[];
    isSelect?: boolean;
    showEdit?: boolean;
    showPreview?: boolean;
    showSelect?: boolean;
    onEdit?: (item: any) => void;
    onPreview?: (item: any) => void;
    idKey?: string;
    onSelectedIdsChange?: (ids: any[]) => void;
    onSelectedDataChange?: (data: any[]) => void;
    isRevisi?: boolean;
}

export default function TableContent({
    columns,
    data,
    isSelect = false,
    showEdit = true,
    showPreview = true,
    showSelect = false,
    onEdit,
    onPreview,
    idKey = 'id',
    onSelectedIdsChange,
    onSelectedDataChange,
    isRevisi = false,
}: TableContentProps) {
    const [selectedIds, setSelectedIds] = useState<any[]>([]);
    const [currentPage, setCurrentPage] = useState(1);
    const [pageSize, setPageSize] = useState(15);

    const totalPages = Math.ceil(data.length / pageSize);
    const startIndex = (currentPage - 1) * pageSize;
    const endIndex = startIndex + pageSize;
    const currentData = data.slice(startIndex, endIndex);

    const handleSelectAll = (checked: boolean) => {
        if (checked) {
            const allIds = currentData.map(item => item[idKey]);
            setSelectedIds(allIds);
            onSelectedIdsChange?.(allIds);
            onSelectedDataChange?.(currentData);
        } else {
            setSelectedIds([]);
            onSelectedIdsChange?.([]);
            onSelectedDataChange?.([]);
        }
    };

    const handleSelectItem = (item: any, checked: boolean) => {
        const itemId = item[idKey];
        const newSelectedIds = checked
            ? [...selectedIds, itemId]
            : selectedIds.filter(id => id !== itemId);

        setSelectedIds(newSelectedIds);
        onSelectedIdsChange?.(newSelectedIds);
        onSelectedDataChange?.(data.filter(d => newSelectedIds.includes(d[idKey])));
    };

    const isSelected = (item: any) => selectedIds.includes(item[idKey]);
    const isAllSelected = currentData.length > 0 && currentData.every(item => isSelected(item));

    const handlePageSizeChange = (size: number) => {
        setPageSize(size);
        setCurrentPage(1);
    };

    const getPageNumbers = () => {
        const pages: (number | '...')[] = [];
        if (totalPages <= 5) {
            for (let i = 1; i <= totalPages; i++) pages.push(i);
        } else {
            pages.push(1);
            if (currentPage > 3) pages.push('...');
            for (let i = Math.max(2, currentPage - 1); i <= Math.min(totalPages - 1, currentPage + 1); i++) pages.push(i);
            if (currentPage < totalPages - 2) pages.push('...');
            pages.push(totalPages);
        }
        return pages;
    };

    return (
        <div
            className="w-full"
            data-aos="fade-up"
            data-aos-duration="600"
        >
            {/* Table Card */}
            <div className="bg-white rounded-2xl shadow-lg border border-gray-100 overflow-hidden">

                {/* Table Header Bar */}
                <div
                    className="px-4 sm:px-6 py-3 sm:py-4 bg-linear-to-r from-primary/8 to-primary/4 border-b border-primary/15 flex items-center justify-between"
                    data-aos="fade-down"
                    data-aos-duration="500"
                    data-aos-delay="100"
                >
                    <div className="flex items-center gap-2.5">
                        <div className="w-7 h-7 bg-linear-to-br from-primary to-primary/70 rounded-lg flex items-center justify-center shadow-sm">
                            <LayoutList className="h-3.5 w-3.5 text-white" />
                        </div>
                        <div>
                            <p className="font-poppins-semibold text-sm text-gray-800">Tabel Data</p>
                            <p className="font-poppins-regular text-xs text-gray-500">{data.length} total data</p>
                        </div>
                    </div>
                    {selectedIds.length > 0 && (
                        <div className="flex items-center gap-2 px-3 py-1.5 bg-primary/10 border border-primary/20 rounded-full animate-pulse">
                            <div className="w-2 h-2 rounded-full bg-primary"></div>
                            <span className="font-poppins-semibold text-xs text-primary">{selectedIds.length} dipilih</span>
                        </div>
                    )}
                </div>

                <div className="overflow-x-auto">
                    <table className="w-full border-collapse">
                        <thead>
                            <tr className="bg-linear-to-r from-primary/12 via-primary/8 to-primary/5 border-b border-primary/20">
                                {isSelect && (
                                    <th className="px-3 sm:px-5 py-3 sm:py-4 text-center w-10">
                                        <div className="flex items-center justify-center">
                                            <input
                                                type="checkbox"
                                                checked={isAllSelected}
                                                onChange={(e) => handleSelectAll(e.target.checked)}
                                                className="w-4 h-4 text-primary bg-white border-2 border-primary/40 rounded-md focus:ring-primary focus:ring-2 cursor-pointer accent-primary transition-all duration-200"
                                            />
                                        </div>
                                    </th>
                                )}
                                {columns.map((column, idx) => (
                                    <th
                                        key={column.key}
                                        className="px-3 sm:px-6 py-3 sm:py-4 text-center font-poppins-semibold text-xs text-gray-700 uppercase tracking-wider whitespace-nowrap"
                                        data-aos="fade-down"
                                        data-aos-duration="400"
                                        data-aos-delay={`${(idx + 1) * 60}`}
                                    >
                                        <span className="inline-flex items-center gap-1.5">
                                            {column.label}
                                        </span>
                                    </th>
                                ))}
                                {(showEdit || showPreview || showSelect) && (
                                    <th
                                        className="px-3 sm:px-5 py-3 sm:py-4 text-center font-poppins-semibold text-xs text-gray-700 uppercase tracking-wider"
                                        data-aos="fade-down"
                                        data-aos-duration="400"
                                        data-aos-delay={`${(columns.length + 1) * 60}`}
                                    >
                                        Aksi
                                    </th>
                                )}
                            </tr>
                        </thead>
                        <tbody className="divide-y divide-gray-100">
                            {currentData.length === 0 ? (
                                <tr>
                                    <td
                                        colSpan={columns.length + (isSelect ? 1 : 0) + (showEdit || showPreview || showSelect ? 1 : 0)}
                                        className="px-4 py-16 text-center"
                                    >
                                        <div
                                            className="flex flex-col items-center gap-4"
                                            data-aos="zoom-in"
                                            data-aos-duration="600"
                                        >
                                            <div className="w-16 h-16 bg-gray-100 rounded-2xl flex items-center justify-center">
                                                <FileX className="w-8 h-8 text-gray-300" />
                                            </div>
                                            <div>
                                                <p className="font-poppins-semibold text-gray-500 text-sm sm:text-base">Tidak ada data</p>
                                                <p className="font-poppins-regular text-gray-400 text-xs sm:text-sm mt-1">Data akan muncul di sini</p>
                                            </div>
                                        </div>
                                    </td>
                                </tr>
                            ) : (
                                currentData.map((item, index) => (
                                    <tr
                                        key={index}
                                        className={`group transition-all duration-200 hover:bg-primary/4 hover:shadow-sm ${isSelected(item) ? 'bg-primary/6 border-l-2 border-l-primary' : 'border-l-2 border-l-transparent'}`}
                                    >
                                        {isSelect && (
                                            <td className="px-3 sm:px-5 py-3 sm:py-4 text-center">
                                                <div className="flex items-center justify-center">
                                                    <input
                                                        type="checkbox"
                                                        checked={isSelected(item)}
                                                        onChange={(e) => handleSelectItem(item, e.target.checked)}
                                                        className="w-4 h-4 text-primary bg-white border-2 border-primary/30 rounded-md focus:ring-primary focus:ring-2 cursor-pointer accent-primary transition-all duration-200"
                                                    />
                                                </div>
                                            </td>
                                        )}
                                        {columns.map((column) => (
                                            <td
                                                key={column.key}
                                                className="px-3 sm:px-6 py-3 sm:py-4 font-poppins-regular text-xs sm:text-sm text-gray-700 text-center"
                                            >
                                                {column.key === 'id' ? (
                                                    <span className="inline-flex items-center justify-center w-6 h-6 bg-primary/10 text-primary font-poppins-semibold text-xs rounded-md">
                                                        {startIndex + index + 1}
                                                    </span>
                                                ) : (
                                                    <span className="group-hover:text-gray-900 transition-colors duration-150">
                                                        {item[column.key]}
                                                    </span>
                                                )}
                                            </td>
                                        ))}
                                        {(showEdit || showPreview || showSelect) && (
                                            <td className="px-3 sm:px-5 py-2.5 sm:py-4">
                                                <div className="flex items-center justify-center gap-1.5 sm:gap-2 flex-wrap">
                                                    {showEdit && (
                                                        <button
                                                            onClick={() => onEdit?.(item)}
                                                            className="group/btn inline-flex items-center gap-1 sm:gap-1.5 px-2.5 sm:px-3.5 py-1.5 sm:py-2 text-primary bg-primary/8 hover:bg-primary hover:text-white rounded-lg transition-all duration-200 cursor-pointer font-poppins-medium text-xs active:scale-95 hover:shadow-md border border-primary/15 hover:border-primary"
                                                            title={isRevisi ? "Revisi" : "Ubah"}
                                                        >
                                                            <Edit2 className="h-3 w-3 sm:h-3.5 sm:w-3.5 transition-transform duration-200 group-hover/btn:rotate-12" />
                                                            <span className="hidden sm:inline">{isRevisi ? "Revisi" : "Ubah"}</span>
                                                        </button>
                                                    )}
                                                    {showPreview && (
                                                        <button
                                                            onClick={() => onPreview?.(item)}
                                                            className="group/btn inline-flex items-center gap-1 sm:gap-1.5 px-2.5 sm:px-3.5 py-1.5 sm:py-2 text-blue-600 bg-blue-50 hover:bg-blue-600 hover:text-white rounded-lg transition-all duration-200 cursor-pointer font-poppins-medium text-xs active:scale-95 hover:shadow-md border border-blue-100 hover:border-blue-600"
                                                            title="Lihat"
                                                        >
                                                            <Eye className="h-3 w-3 sm:h-3.5 sm:w-3.5 transition-transform duration-200 group-hover/btn:scale-110" />
                                                            <span className="hidden sm:inline">Lihat</span>
                                                        </button>
                                                    )}
                                                    {showSelect && (
                                                        <button
                                                            onClick={() => onSelectedDataChange?.(item)}
                                                            className="group/btn inline-flex items-center gap-1 sm:gap-1.5 px-2.5 sm:px-3.5 py-1.5 sm:py-2 text-green-600 bg-green-50 hover:bg-green-600 hover:text-white rounded-lg transition-all duration-200 cursor-pointer font-poppins-medium text-xs active:scale-95 hover:shadow-md border border-green-100 hover:border-green-600"
                                                            title="Pilih"
                                                        >
                                                            <CheckCircle className="h-3 w-3 sm:h-3.5 sm:w-3.5 transition-transform duration-200 group-hover/btn:scale-110" />
                                                            <span className="hidden sm:inline">Pilih</span>
                                                        </button>
                                                    )}
                                                </div>
                                            </td>
                                        )}
                                    </tr>
                                ))
                            )}
                        </tbody>
                    </table>
                </div>

                {data.length > 0 && (
                    <div
                        className="px-4 sm:px-6 py-3 sm:py-4 border-t border-gray-100 bg-linear-to-r from-gray-50/80 to-white flex flex-col gap-3 sm:flex-row sm:items-center sm:justify-between"
                    >
                        <div className="flex items-center gap-2">
                            <span className="font-poppins-regular text-xs text-gray-500 whitespace-nowrap">Tampilkan:</span>
                            <select
                                value={pageSize}
                                onChange={(e) => handlePageSizeChange(Number(e.target.value))}
                                className="px-2.5 py-1.5 border border-gray-200 rounded-lg font-poppins-medium text-xs text-gray-700 focus:outline-none focus:border-primary focus:ring-2 focus:ring-primary/20 bg-white cursor-pointer transition-all duration-200 hover:border-primary/40"
                            >
                                {[15, 30, 50, 100].map(n => <option key={n} value={n}>{n}</option>)}
                            </select>
                            <span className="font-poppins-regular text-xs text-gray-500">data</span>
                        </div>

                        {/* Info */}
                        <div className="flex items-center gap-1.5">
                            <span className="font-poppins-regular text-xs text-gray-500">
                                Menampilkan
                            </span>
                            <span className="font-poppins-semibold text-xs text-gray-700 px-2 py-0.5 bg-primary/8 rounded-md">
                                {startIndex + 1}–{Math.min(endIndex, data.length)}
                            </span>
                            <span className="font-poppins-regular text-xs text-gray-500">
                                dari <span className="font-poppins-semibold text-gray-700">{data.length}</span> data
                            </span>
                        </div>

                        {/* Page buttons */}
                        <div className="flex items-center gap-1.5">
                            <button
                                onClick={() => setCurrentPage(p => Math.max(1, p - 1))}
                                disabled={currentPage === 1}
                                className="inline-flex items-center gap-1 px-2.5 sm:px-3 py-1.5 sm:py-2 border border-gray-200 rounded-lg font-poppins-medium text-xs text-gray-600 hover:bg-primary/8 hover:border-primary/30 hover:text-primary disabled:opacity-40 disabled:cursor-not-allowed transition-all duration-200 active:scale-95"
                            >
                                <ChevronLeft className="h-3.5 w-3.5" />
                                <span className="hidden sm:inline">Prev</span>
                            </button>

                            <div className="flex items-center gap-1">
                                {getPageNumbers().map((page, idx) =>
                                    page === '...' ? (
                                        <span key={`ellipsis-${idx}`} className="px-2 py-1 text-xs text-gray-400 font-poppins-regular">…</span>
                                    ) : (
                                        <button
                                            key={page}
                                            onClick={() => setCurrentPage(page as number)}
                                            className={`w-7 h-7 sm:w-8 sm:h-8 rounded-lg font-poppins-semibold text-xs transition-all duration-200 active:scale-95 ${
                                                currentPage === page
                                                    ? 'bg-primary text-white shadow-md shadow-primary/30 scale-105'
                                                    : 'border border-gray-200 text-gray-600 hover:bg-primary/8 hover:border-primary/30 hover:text-primary'
                                            }`}
                                        >
                                            {page}
                                        </button>
                                    )
                                )}
                            </div>

                            <button
                                onClick={() => setCurrentPage(p => Math.min(totalPages, p + 1))}
                                disabled={currentPage === totalPages}
                                className="inline-flex items-center gap-1 px-2.5 sm:px-3 py-1.5 sm:py-2 border border-gray-200 rounded-lg font-poppins-medium text-xs text-gray-600 hover:bg-primary/8 hover:border-primary/30 hover:text-primary disabled:opacity-40 disabled:cursor-not-allowed transition-all duration-200 active:scale-95"
                            >
                                <span className="hidden sm:inline">Next</span>
                                <ChevronRight className="h-3.5 w-3.5" />
                            </button>
                        </div>
                    </div>
                )}
            </div>
        </div>
    );
}