<?xml version="1.0" encoding="utf-8"?>
<xsl:stylesheet version="1.0"
      xmlns:xsl="http://www.w3.org/1999/XSL/Transform"
      xmlns:fo="http://www.w3.org/1999/XSL/Format">
  <xsl:output method="xml" indent="yes"/>
  <xsl:template match="/">
    <fo:root>
      <fo:layout-master-set>
        <fo:simple-page-master master-name="A4-portrait"
              page-height="29.7cm" page-width="21.0cm" margin="2.5cm">
          <fo:region-body/>
        </fo:simple-page-master>
      </fo:layout-master-set>
      <fo:page-sequence master-reference="A4-portrait">
        <fo:flow flow-name="xsl-region-body">
            <fo:block space-after="5em">
                <fo:external-graphic src="images/ons-logo.svg" height="2em" />
            </fo:block>
            <fo:block font-family="Hind" font-size="12pt" space-after="1em" font-weight="bold">
                Welcome to Digital Publishing!
            </fo:block>
            <xsl:for-each select="welcome-pack/sections/section">
                <xsl:if test="title">
                    <fo:block font-family="Hind" font-size="12pt" space-after="1em" font-weight="bold">
                        <xsl:value-of select="title"/>
                    </fo:block>
                </xsl:if>
                <xsl:for-each select="p">
                    <fo:block font-family="Hind" font-size="12pt" space-after="1em" font-weight="normal">
                        <xsl:value-of select="."/>
                    </fo:block>
                </xsl:for-each>
            </xsl:for-each>
        </fo:flow>
      </fo:page-sequence>
    </fo:root>
  </xsl:template>
</xsl:stylesheet>