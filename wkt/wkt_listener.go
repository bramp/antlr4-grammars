// Code generated from wkt.g4 by ANTLR 4.9.3. DO NOT EDIT.

package wkt // wkt
import "github.com/antlr/antlr4/runtime/Go/antlr"

// wktListener is a complete listener for a parse tree produced by wktParser.
type wktListener interface {
	antlr.ParseTreeListener

	// EnterGeometryCollection is called when entering the geometryCollection production.
	EnterGeometryCollection(c *GeometryCollectionContext)

	// EnterGeometry is called when entering the geometry production.
	EnterGeometry(c *GeometryContext)

	// EnterPointGeometry is called when entering the pointGeometry production.
	EnterPointGeometry(c *PointGeometryContext)

	// EnterLineStringGeometry is called when entering the lineStringGeometry production.
	EnterLineStringGeometry(c *LineStringGeometryContext)

	// EnterPolygonGeometry is called when entering the polygonGeometry production.
	EnterPolygonGeometry(c *PolygonGeometryContext)

	// EnterMultiCurveGeometry is called when entering the multiCurveGeometry production.
	EnterMultiCurveGeometry(c *MultiCurveGeometryContext)

	// EnterMultiSurfaceGeometry is called when entering the multiSurfaceGeometry production.
	EnterMultiSurfaceGeometry(c *MultiSurfaceGeometryContext)

	// EnterCurvePolygonGeometry is called when entering the curvePolygonGeometry production.
	EnterCurvePolygonGeometry(c *CurvePolygonGeometryContext)

	// EnterCompoundCurveGeometry is called when entering the compoundCurveGeometry production.
	EnterCompoundCurveGeometry(c *CompoundCurveGeometryContext)

	// EnterMultiPointGeometry is called when entering the multiPointGeometry production.
	EnterMultiPointGeometry(c *MultiPointGeometryContext)

	// EnterMultiLineStringGeometry is called when entering the multiLineStringGeometry production.
	EnterMultiLineStringGeometry(c *MultiLineStringGeometryContext)

	// EnterMultiPolygonGeometry is called when entering the multiPolygonGeometry production.
	EnterMultiPolygonGeometry(c *MultiPolygonGeometryContext)

	// EnterMultiPolyhedralSurfaceGeometry is called when entering the multiPolyhedralSurfaceGeometry production.
	EnterMultiPolyhedralSurfaceGeometry(c *MultiPolyhedralSurfaceGeometryContext)

	// EnterMultiTinGeometry is called when entering the multiTinGeometry production.
	EnterMultiTinGeometry(c *MultiTinGeometryContext)

	// EnterCircularStringGeometry is called when entering the circularStringGeometry production.
	EnterCircularStringGeometry(c *CircularStringGeometryContext)

	// EnterPointOrClosedPoint is called when entering the pointOrClosedPoint production.
	EnterPointOrClosedPoint(c *PointOrClosedPointContext)

	// EnterPolygon is called when entering the polygon production.
	EnterPolygon(c *PolygonContext)

	// EnterLineString is called when entering the lineString production.
	EnterLineString(c *LineStringContext)

	// EnterPoint is called when entering the point production.
	EnterPoint(c *PointContext)

	// EnterName is called when entering the name production.
	EnterName(c *NameContext)

	// ExitGeometryCollection is called when exiting the geometryCollection production.
	ExitGeometryCollection(c *GeometryCollectionContext)

	// ExitGeometry is called when exiting the geometry production.
	ExitGeometry(c *GeometryContext)

	// ExitPointGeometry is called when exiting the pointGeometry production.
	ExitPointGeometry(c *PointGeometryContext)

	// ExitLineStringGeometry is called when exiting the lineStringGeometry production.
	ExitLineStringGeometry(c *LineStringGeometryContext)

	// ExitPolygonGeometry is called when exiting the polygonGeometry production.
	ExitPolygonGeometry(c *PolygonGeometryContext)

	// ExitMultiCurveGeometry is called when exiting the multiCurveGeometry production.
	ExitMultiCurveGeometry(c *MultiCurveGeometryContext)

	// ExitMultiSurfaceGeometry is called when exiting the multiSurfaceGeometry production.
	ExitMultiSurfaceGeometry(c *MultiSurfaceGeometryContext)

	// ExitCurvePolygonGeometry is called when exiting the curvePolygonGeometry production.
	ExitCurvePolygonGeometry(c *CurvePolygonGeometryContext)

	// ExitCompoundCurveGeometry is called when exiting the compoundCurveGeometry production.
	ExitCompoundCurveGeometry(c *CompoundCurveGeometryContext)

	// ExitMultiPointGeometry is called when exiting the multiPointGeometry production.
	ExitMultiPointGeometry(c *MultiPointGeometryContext)

	// ExitMultiLineStringGeometry is called when exiting the multiLineStringGeometry production.
	ExitMultiLineStringGeometry(c *MultiLineStringGeometryContext)

	// ExitMultiPolygonGeometry is called when exiting the multiPolygonGeometry production.
	ExitMultiPolygonGeometry(c *MultiPolygonGeometryContext)

	// ExitMultiPolyhedralSurfaceGeometry is called when exiting the multiPolyhedralSurfaceGeometry production.
	ExitMultiPolyhedralSurfaceGeometry(c *MultiPolyhedralSurfaceGeometryContext)

	// ExitMultiTinGeometry is called when exiting the multiTinGeometry production.
	ExitMultiTinGeometry(c *MultiTinGeometryContext)

	// ExitCircularStringGeometry is called when exiting the circularStringGeometry production.
	ExitCircularStringGeometry(c *CircularStringGeometryContext)

	// ExitPointOrClosedPoint is called when exiting the pointOrClosedPoint production.
	ExitPointOrClosedPoint(c *PointOrClosedPointContext)

	// ExitPolygon is called when exiting the polygon production.
	ExitPolygon(c *PolygonContext)

	// ExitLineString is called when exiting the lineString production.
	ExitLineString(c *LineStringContext)

	// ExitPoint is called when exiting the point production.
	ExitPoint(c *PointContext)

	// ExitName is called when exiting the name production.
	ExitName(c *NameContext)
}
