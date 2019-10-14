package pg

func (h *Handler) clearRankingCache(eventID int64) {
	h.rankingCacheMu.Lock()
	defer h.rankingCacheMu.Unlock()
	delete(h.rankingCache, eventID)
}

func (h *Handler) clearResultCache(eventID int64) {
	h.resultCacheMu.Lock()
	defer h.resultCacheMu.Unlock()
	delete(h.resultCache, eventID)
}
